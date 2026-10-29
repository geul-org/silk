//ff:func feature=index type=rule control=sequence
//ff:what Verify — the catalog matches the archive, lists exactly the maps it needs, and each map matches its pack; the number checked
package index

import (
	"context"
	"fmt"
	"slices"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// Verify checks the catalog (checkSets) and each target's map (verifyOne).
func Verify(ctx context.Context, a *archive.Archive, dir string) (int, error) {
	c, err := readCatalog(dir)
	if err != nil {
		return 0, err
	}
	if err := checkSets(a, c); err != nil {
		return 0, err
	}
	var rels []string
	err = eachTarget(ctx, a, func(t Target) error {
		rels = append(rels, mapRel(t))
		return verifyOne(mapPath(dir, mapRel(t)), t)
	})
	if err != nil {
		return 0, err
	}
	if !slices.Equal(rels, c.Maps) {
		return 0, fmt.Errorf("%w: catalog lists %d maps, archive needs %d", ErrStale, len(c.Maps), len(rels))
	}
	return len(rels), nil
}
