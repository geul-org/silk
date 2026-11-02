//ff:func feature=index type=scanner control=sequence
//ff:what TestDocPrint — dead numbers do not count; rows print with paths
package index

import (
	"context"
	"testing"
)

func TestDocPrint(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 3)
	ix, _ := Open(ctx, a, dir, 0)
	defer ix.Close()
	p := ix.docPrint()
	if p.n != 3 {
		t.Fatal(p)
	}
	s := ix.secs["t"][0]
	r := ix.rowPrint(s)
	if r.n != 3 {
		t.Fatal(r)
	}
	ix.Live.Put(1, false)
	if ix.docPrint().n != 2 {
		t.Fatal("dead")
	}
	var f fingerprint
	k, reg := s.Files[0].PartAt(0)
	ix.printRows(&f, s.Files[0], k, reg)
	if f.n != 2 {
		t.Fatal(f)
	}
}
