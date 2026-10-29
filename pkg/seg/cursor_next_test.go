//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what TestCursorNext — rows over chunk boundaries with the partition slot stripped, then false; a read error is reported
package seg

import (
	"path/filepath"
	"testing"
)

func TestCursorNext(t *testing.T) {
	b := NewBucket(Info{PartWidth: 2, SortWidth: 2}, 1)
	for i := range chunkRows + 3 {
		b.Add([]byte("aa"), row("k1", uint32(i), 0))
	}
	p := filepath.Join(t.TempDir(), "c.sseg")
	b.Write(p, true)
	f, _ := Open(p)
	defer f.Close()
	c := &cursor{g: Group{f: f, rg: f.partRegion(0)}}
	n := 0
	for {
		r, ok, err := c.next()
		if err != nil || (ok && len(r) != 10) {
			t.Fatal(err, r)
		}
		if !ok {
			break
		}
		n++
	}
	if n != chunkRows+3 {
		t.Fatal(n)
	}
	o := sampleSeg(t, true)
	c2 := &cursor{g: Group{f: o, rg: o.otherRegion(), strip: 2}}
	if r, _, _ := c2.next(); string(r[:2]) != "y1" {
		t.Fatal(string(r))
	}
	c3 := &cursor{g: Group{f: failing(f, 0), rg: f.partRegion(0)}}
	if _, _, err := c3.next(); err == nil {
		t.Fatal("read error")
	}
}
