//ff:func feature=index type=reader control=sequence
//ff:what TestPackPaths — `<set dir>/o/<hash>.gpk` per pack; an unknown set is ErrStale
package index

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestPackPaths(t *testing.T) {
	a, _, _ := built(t)
	ps, err := packPaths(a, []PackRef{{Set: a.Sets[0].Name, Hash: "ab"}})
	if err != nil || ps[0] != filepath.Join(a.Sets[0].Dir, "o", "ab.gpk") {
		t.Fatal(ps, err)
	}
	if _, err := packPaths(a, []PackRef{{Set: "0x9999"}}); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
}
