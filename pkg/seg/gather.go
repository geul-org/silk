//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what gather — read every row of the groups (an unsorted input among them), sort them and add the live ones
package seg

// gather is for appended segments, which are small.
func gather(gs []Group, w *Writer, live func(uint32) bool) (uint64, error) {
	var all []byte
	for _, g := range gs {
		b, err := readGroup(g)
		if err != nil {
			return 0, err
		}
		all = append(all, b...)
	}
	return addRows(w, SortRows(all, gs[0].rg.w-gs[0].strip), live)
}
