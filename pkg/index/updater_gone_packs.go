//ff:func feature=index type=helper control=iteration dimension=1
//ff:what updater.gonePacks — the pack numbers the current sets no longer list
package index

func (u *updater) gonePacks() map[uint32]bool {
	out := map[uint32]bool{}
	for i, p := range u.m.Packs {
		s := u.ar.Set(p.Set)
		if s == nil {
			out[uint32(i)] = true
			continue
		}
		if _, ok := s.ObjectOf(p.Hash); !ok {
			out[uint32(i)] = true
		}
	}
	return out
}
