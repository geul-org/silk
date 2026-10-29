//ff:func feature=recmap type=reader control=sequence
//ff:what File.Offsets — the ascending pack offsets of one key's records, one ReadAt of its off[] range
package recmap

// Offsets reads key's slice of off[]; a key the map lacks has none.
func (f *File) Offsets(key uint16) ([]uint32, error) {
	k, ok := f.Range(key)
	if !ok {
		return nil, nil
	}
	return f.offs(k.Start, k.Count)
}
