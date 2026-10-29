//ff:func feature=seg type=writer control=iteration dimension=1
//ff:what addRows — add the rows of b (the Writer's row width each) to the open partition, dropping dead ones; the number dropped
package seg

// addRows walks b in rows of w bytes (the Writer's partition row width).
func addRows(w *Writer, b []byte, live func(uint32) bool) (uint64, error) {
	r, _ := w.info.RowWidth()
	dropped := uint64(0)
	for ; len(b) >= r; b = b[r:] {
		d, err := emitRow(w, b[:r], live)
		dropped += d
		if err != nil {
			return dropped, err
		}
	}
	return dropped, nil
}
