//ff:func feature=archive type=reader control=iteration dimension=1
//ff:what Set.Data — the data packs of the set in replay order
package archive

func (s *Set) Data() []Data {
	var out []Data
	for _, o := range s.r.DataPacks() {
		out = append(out, Data{Obj: o, Hash: s.Objects[o].Hash})
	}
	return out
}
