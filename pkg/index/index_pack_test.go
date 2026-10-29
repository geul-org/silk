//ff:func feature=index type=reader control=sequence
//ff:what TestIndexPack — a data pack opened once and reused; a missing file fails
package index

import "testing"

func TestIndexPack(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	f1, err := ix.pack(0)
	f2, _ := ix.pack(0)
	if err != nil || f1 != f2 {
		t.Fatal(err)
	}
	ix.packs[0], ix.paths[0] = nil, ix.paths[0]+".none"
	if _, err := ix.pack(0); err == nil {
		t.Fatal("missing pack")
	}
}
