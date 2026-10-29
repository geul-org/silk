//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what scanRows — the fused loop over rows [lo, hi): test each row against the pattern and keep its number
package sig

// scanRows tests rows lo…hi-1 and builds no intermediate arrays.
func scanRows(t *Table, p Pattern, lo, hi int) []uint32 {
	var out []uint32
	w := t.Words
	for i := lo; i < hi; i++ {
		if p.match(t.Data[i*w : (i+1)*w]) {
			out = append(out, uint32(i))
		}
	}
	return out
}
