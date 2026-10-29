//ff:func feature=index type=builder control=sequence
//ff:what Build — write the record map of every data pack and Issuer of an archive, then the catalog, then drop maps no longer listed
package index

import (
	"context"
	"os"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/recmap"
)

// Build makes each map (buildOne) and returns their headers in build order.
// The catalog is written after every map, so a failed build leaves the
// previous catalog (whose maps may now be gone: Open then fails, not lies).
func Build(ctx context.Context, a *archive.Archive, dir string) ([]recmap.Header, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	var out []recmap.Header
	var rels []string
	err := eachTarget(ctx, a, func(t Target) error {
		h, err := buildOne(mapPath(dir, mapRel(t)), t)
		out = append(out, h)
		rels = append(rels, mapRel(t))
		return err
	})
	if err != nil {
		return nil, err
	}
	if err := writeCatalog(dir, newCatalog(a, rels)); err != nil {
		return nil, err
	}
	return out, removeStale(dir, rels)
}
