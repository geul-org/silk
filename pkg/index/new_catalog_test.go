//ff:func feature=index type=builder control=sequence
//ff:what TestNewCatalog — version, sets in archive order and the given maps
package index

import (
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestNewCatalog(t *testing.T) {
	a := &archive.Archive{Sets: []archive.Set{{Name: "0x0001", Generation: 2, Manifest: "m", Dir: "/x"}}}
	c := newCatalog(a, []string{"0x0001/p.srm"})
	if c.Version != catalogVersion || len(c.Sets) != 1 || c.Sets[0] != (CatalogSet{Name: "0x0001", Generation: 2, Manifest: "m"}) || c.Maps[0] != "0x0001/p.srm" {
		t.Fatal(c)
	}
}
