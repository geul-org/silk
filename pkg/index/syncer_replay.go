//ff:func feature=index type=reader control=iteration dimension=1
//ff:what syncer.replay — walk the sets and their unseen data packs
package index

import "context"

func (s *syncer) replay(ctx context.Context) error {
	seen := map[PackRef]bool{}
	for _, p := range s.u.m.Packs {
		seen[p] = true
	}
	for _, st := range s.u.ar.Sets {
		if err := s.replaySet(ctx, st, seen); err != nil {
			return err
		}
	}
	return nil
}
