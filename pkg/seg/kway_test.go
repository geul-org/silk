//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what TestKway — two sorted groups interleave in order with dead rows dropped; read and write errors stop the merge
package seg

import (
	"path/filepath"
	"testing"
)

func TestKway(t *testing.T) {
	f := sampleSeg(t, true)
	gs, _ := f.Groups()
	p := filepath.Join(t.TempDir(), "k.sseg")
	w, _ := NewWriter(p, Info{PartWidth: 2, SortWidth: 2, Sorted: true}, 1)
	w.Begin([]byte("aa"))
	d, err := kway([]Group{gs[0], gs[1]}, w, func(doc uint32) bool { return doc != 4 })
	w.End()
	w.Finish()
	m, _ := Open(p)
	defer m.Close()
	if err != nil || d != 1 || m.Info.Rows != 4 || m.Verify() != nil {
		t.Fatal(err, d, m.Info)
	}
	bad := Group{f: failing(f, 0), rg: gs[0].rg}
	if _, err := kway([]Group{bad}, w, allLive); err == nil {
		t.Fatal("first read")
	}
	late := Group{f: failing(f, 1), rg: gs[0].rg}
	w2, _ := NewWriter(filepath.Join(t.TempDir(), "x.sseg"), Info{PartWidth: 2, SortWidth: 2}, 0)
	defer w2.Abort()
	w2.Begin([]byte("aa"))
	big := NewBucket(Info{PartWidth: 2, SortWidth: 2}, 1)
	for i := range chunkRows + 1 {
		big.Add([]byte("aa"), row("k1", uint32(i), 0))
	}
	bp := filepath.Join(t.TempDir(), "big.sseg")
	big.Write(bp, true)
	bf, _ := Open(bp)
	defer bf.Close()
	late.f, late.rg = failing(bf, 1), bf.partRegion(0)
	if _, err := kway([]Group{late}, w2, allLive); err == nil {
		t.Fatal("second chunk")
	}
	if _, err := kway([]Group{gs[0]}, brokenAt(t), allLive); err == nil {
		t.Fatal("write error")
	}
}
