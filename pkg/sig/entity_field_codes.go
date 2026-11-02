//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what Entity.fieldCodes — the codes of each field
package sig

func (e *Entity) fieldCodes(fs []*Field) [][]uint64 {
	out := make([][]uint64, len(fs))
	for i, f := range fs {
		out[i] = e.codes(f)
	}
	return out
}
