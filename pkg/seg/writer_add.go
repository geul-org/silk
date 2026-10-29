//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Add — one row of the open partition: buffered while the partition is below Min, then written straight through
package seg

// Add takes a row of SortWidth+8 bytes; the caller orders rows in a sorted
// segment.
func (w *Writer) Add(row []byte) error {
	w.n++
	if w.n < uint64(w.Min) {
		w.buf = append(w.buf, row...)
		return nil
	}
	if len(w.buf) > 0 {
		if err := w.write(w.buf); err != nil {
			return err
		}
		w.buf = w.buf[:0]
	}
	return w.write(row)
}
