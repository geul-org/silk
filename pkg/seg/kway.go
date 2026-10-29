//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what kway — streaming k-way merge of sorted groups into the open partition, smallest row first, dead rows dropped
package seg

// kway reads each group once; O(n·k) comparisons with small k.
func kway(gs []Group, w *Writer, live func(uint32) bool) (uint64, error) {
	cs := make([]*cursor, len(gs))
	heads := make([][]byte, len(gs))
	for i, g := range gs {
		cs[i] = &cursor{g: g}
		var err error
		if heads[i], err = advance(cs[i]); err != nil {
			return 0, err
		}
	}
	dropped := uint64(0)
	for k := minHead(heads); k >= 0; k = minHead(heads) {
		d, err := emitRow(w, heads[k], live)
		dropped += d
		if err != nil {
			return dropped, err
		}
		if heads[k], err = advance(cs[k]); err != nil {
			return dropped, err
		}
	}
	return dropped, nil
}
