//ff:func feature=seg type=reader control=sequence
//ff:what TestFileOtherGroups — consecutive other rows of one key form one group with its own region; a read error fails
package seg

import (
	"path/filepath"
	"testing"
)

func TestFileOtherGroups(t *testing.T) {
	b := NewBucket(Info{PartWidth: 2, SortWidth: 2}, 3)
	b.Add([]byte("aa"), row("k1", 1, 0))
	b.Add([]byte("aa"), row("k2", 1, 2))
	b.Add([]byte("bb"), row("k1", 2, 0))
	p := filepath.Join(t.TempDir(), "o.sseg")
	b.Write(p, true)
	f, _ := Open(p)
	defer f.Close()
	gs, err := f.otherGroups(nil)
	if err != nil || len(gs) != 2 || gs[0].rg.n != 2 || gs[1].rg.n != 1 || gs[1].rg.off != gs[0].rg.off+24 {
		t.Fatal(err, gs)
	}
	if _, err := failing(f, 0).otherGroups(nil); err == nil {
		t.Fatal("read error")
	}
}
