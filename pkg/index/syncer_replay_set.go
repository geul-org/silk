//ff:func feature=index type=reader control=iteration dimension=1
//ff:what syncer.replaySet — one set's unseen data packs, oldest first
package index

import (
	"context"

	"github.com/geul-org/silk/pkg/archive"
)

func (s *syncer) replaySet(ctx context.Context, st *archive.Set, seen map[PackRef]bool) error {
	for _, d := range st.Data() {
		ref := PackRef{Set: st.Name, Hash: d.Hash}
		if seen[ref] {
			continue
		}
		s.pack = uint32(len(s.u.m.Packs))
		s.u.m.Packs = append(s.u.m.Packs, ref)
		if err := st.Steps(ctx, d.Obj, s.step); err != nil {
			return err
		}
	}
	return nil
}
