//ff:func feature=index type=rule control=sequence
//ff:what sameLoc — an indexed version lies at the same set, data pack, offset and length as an archive document
package index

// sameLoc means the bytes did not move (pack objects are immutable).
func sameLoc(m Manifest, e DocEntry, d liveDoc) bool {
	p := m.Packs[e.Pack]
	return p.Set == d.Set && p.Hash == d.Pack && e.Off == d.Doc.Off && uint64(e.Len) == d.Doc.Len
}
