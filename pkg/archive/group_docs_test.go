//ff:func feature=archive type=builder control=sequence
//ff:what TestGroupDocs — version paths grouped by object with their places; other paths skipped; zstd and out-of-range objects rejected
package archive

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestGroupDocs(t *testing.T) {
	v := "0x00010000/0x5100000000001d52/0x00000001.geul"
	st := map[string]pack.Loc{
		v: {Obj: 1, Off: 100, Len: 10},
		"0x00010000/0x5100000000001d52/versions.geul": {Obj: 1, Off: 90, Len: 10},
		"issuers.geul": {Obj: 1, Off: 80, Len: 10},
	}
	by, err := groupDocs(st, 2)
	want := Doc{Issuer: "0x00010000", ID: 0x5100000000001d52, Version: 1, Off: 100, Len: 10}
	if err != nil || len(by) != 1 || len(by[1]) != 1 || by[1][0] != want {
		t.Fatal(by, err)
	}
	if _, err := groupDocs(map[string]pack.Loc{v: {Obj: 1, Enc: 1}}, 2); !errors.Is(err, ErrCompressed) {
		t.Fatal(err)
	}
	if _, err := groupDocs(map[string]pack.Loc{v: {Obj: 2}}, 2); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
}
