//ff:func feature=seg type=writer control=iteration dimension=1
//ff:what Writer.End — close the open partition: a table entry when it reached Min rows, else its rows join the other partition behind the key
package seg

// End finishes the partition Begin opened; an empty one leaves no trace.
func (w *Writer) End() {
	w.open = false
	if w.n == 0 {
		return
	}
	if w.n >= uint64(w.Min) {
		w.parts = append(w.parts, Part{Key: w.key, Start: w.info.Rows, Count: w.n})
		w.info.Rows += w.n
		return
	}
	r, _ := w.info.RowWidth()
	for b := w.buf; len(b) > 0; b = b[r:] {
		w.other = append(append(w.other, w.key...), b[:r]...)
	}
	w.info.Other += w.n
}
