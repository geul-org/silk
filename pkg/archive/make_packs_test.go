//ff:func feature=archive type=builder control=sequence
//ff:what TestMakePacks — data objects with documents become Packs in object order with sorted documents; documents in an index object fail
package archive

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestMakePacks(t *testing.T) {
	s := Set{Name: "0x0001", Dir: "/d"}
	objs := []Object{{Kind: 2, Hash: "aa"}, {Kind: 1, Size: 9, Hash: "bb"}, {Kind: 1, Hash: "cc"}}
	by := map[uint32][]Doc{1: {{ID: 2, Off: 70}, {ID: 1, Off: 64}}}
	ps, err := makePacks(s, objs, by)
	if err != nil || len(ps) != 1 || ps[0].Hash != "bb" || ps[0].Size != 9 || ps[0].Path != filepath.Join("/d", "o", "bb.gpk") || ps[0].Docs[0].ID != 1 {
		t.Fatal(ps, err)
	}
	if _, err := makePacks(s, objs, map[uint32][]Doc{0: {{ID: 1}}}); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
}
