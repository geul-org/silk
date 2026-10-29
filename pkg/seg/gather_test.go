//ff:func feature=seg type=builder control=sequence
//ff:what TestGather — rows of appended groups sorted into the open partition, dead rows dropped; a read error fails
package seg

import (
	"path/filepath"
	"testing"
)

func TestGather(t *testing.T) {
	f := sampleSeg(t, false)
	gs, _ := f.Groups()
	p := filepath.Join(t.TempDir(), "g.sseg")
	w, _ := NewWriter(p, Info{PartWidth: 2, SortWidth: 2, Sorted: true}, 0)
	w.Begin([]byte("aa"))
	d, err := gather([]Group{gs[0], gs[2]}, w, func(doc uint32) bool { return doc != 1 })
	w.End()
	w.Finish()
	m, _ := Open(p)
	defer m.Close()
	if err != nil || d != 1 || m.Info.Rows != 3 || m.Verify() != nil {
		t.Fatal(err, d, m.Info)
	}
	if _, err := gather([]Group{{f: failing(f, 0), rg: gs[0].rg}}, w, allLive); err == nil {
		t.Fatal("read error")
	}
}
