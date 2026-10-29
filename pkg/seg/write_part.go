//ff:func feature=seg type=writer control=sequence
//ff:what writePart — one partition of in-memory rows to a Writer (sorted first for a run)
package seg

// writePart begins, adds every row and ends partition key.
func writePart(w *Writer, key, rows []byte, sorted bool) error {
	r, _ := w.info.RowWidth()
	if sorted {
		rows = SortRows(rows, r)
	}
	if err := w.Begin(key); err != nil {
		return err
	}
	_, err := addRows(w, rows, func(uint32) bool { return true })
	w.End()
	return err
}
