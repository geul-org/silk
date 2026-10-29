//ff:func feature=index type=helper control=iteration dimension=1
//ff:what eachTarget — call fn for every (data pack, Issuer) map an archive needs, set by set
package index

import (
	"context"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// eachTarget lists each set's data packs (archive.Set.Packs) and visits them.
func eachTarget(ctx context.Context, a *archive.Archive, fn func(Target) error) error {
	for _, s := range a.Sets {
		packs, err := s.Packs(ctx)
		if err != nil {
			return err
		}
		if err := eachPack(packs, fn); err != nil {
			return err
		}
	}
	return nil
}
