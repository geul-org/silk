//ff:func feature=archive type=reader control=sequence
//ff:what Set.Steps — replay one data pack's inventory, handing fn each version document step in order; other paths are skipped
package archive

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

func (s *Set) Steps(ctx context.Context, obj uint32, fn func(Step) error) error {
	return s.r.PackInventory(ctx, obj, func(e pack.InventoryEntry) error {
		st, ok := parseVersionPath(e.Path)
		if !ok {
			return nil
		}
		st.Loc, st.Deleted = e.Loc, e.Deleted
		return fn(st)
	})
}
