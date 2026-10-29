//ff:func feature=index type=rule control=iteration dimension=1
//ff:what checkSets — the catalog's sets are the archive's current sets (same names, generations, manifests, order)
package index

import (
	"fmt"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// checkSets reports ErrStale when a set was added, removed or got a new generation.
func checkSets(a *archive.Archive, c Catalog) error {
	if len(c.Sets) != len(a.Sets) {
		return fmt.Errorf("%w: %d sets indexed, archive has %d", ErrStale, len(c.Sets), len(a.Sets))
	}
	for i, s := range a.Sets {
		if c.Sets[i] != (CatalogSet{Name: s.Name, Generation: s.Generation, Manifest: s.Manifest}) {
			return fmt.Errorf("%w: set %q is at generation %d (%s)", ErrStale, s.Name, s.Generation, s.Manifest)
		}
	}
	return nil
}
