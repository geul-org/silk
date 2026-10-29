//ff:func feature=spec type=encoder control=sequence
//ff:what Tag.Extract — a record's partition slot then sort slot; a missing field gives a zero slot
package spec

// Extract appends the record's keys to dst (PartWidth + SortWidth bytes).
func (t Tag) Extract(dst, rec []byte) []byte {
	return appendKey(appendKey(dst, rec, t.Partition), rec, t.Sort)
}
