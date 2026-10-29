//ff:func feature=index type=rule control=sequence
//ff:what TestCheckSets — equal set lists pass; another count or generation is ErrStale
package index

import (
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestCheckSets(t *testing.T) {
	a := &archive.Archive{Sets: []archive.Set{{Name: "a", Generation: 1, Manifest: "m"}}}
	c := Catalog{Sets: []CatalogSet{{Name: "a", Generation: 1, Manifest: "m"}}}
	if err := checkSets(a, c); err != nil {
		t.Fatal(err)
	}
	if err := checkSets(&archive.Archive{}, c); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	c.Sets[0].Generation = 2
	if err := checkSets(a, c); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
}
