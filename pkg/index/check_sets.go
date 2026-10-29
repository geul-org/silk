//ff:func feature=index type=rule control=iteration dimension=1
//ff:what checkSets — the indexed set generations are the archive's current ones (names, generations, manifests, order); else ErrStale
package index

import (
	"fmt"

	"github.com/geul-org/silk/pkg/archive"
)

// checkSets compares the manifest's sets with a's.
func checkSets(a *archive.Archive, sets []SetRef) error {
	if len(sets) != len(a.Sets) {
		return fmt.Errorf("%w: %d sets indexed, archive has %d", ErrStale, len(sets), len(a.Sets))
	}
	for i, s := range a.Sets {
		if sets[i] != (SetRef{Name: s.Name, Generation: s.Generation, Manifest: s.Manifest}) {
			return fmt.Errorf("%w: set %q is at generation %d", ErrStale, s.Name, s.Generation)
		}
	}
	return nil
}
