//ff:func feature=archive type=helper control=iteration dimension=1
//ff:what Set.ObjectOf — the object number of a pack hash in this generation; false when the pack is gone
package archive

func (s *Set) ObjectOf(hash string) (uint32, bool) {
	for i, o := range s.Objects {
		if o.Hash == hash {
			return uint32(i), true
		}
	}
	return 0, false
}
