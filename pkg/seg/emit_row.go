//ff:func feature=seg type=writer control=sequence
//ff:what emitRow — add a row to the open partition when its document is live; 1 when dropped
package seg

// emitRow filters by the live predicate.
func emitRow(w *Writer, row []byte, live func(uint32) bool) (uint64, error) {
	doc, _ := SeqOf(row)
	if !live(doc) {
		return 1, nil
	}
	return 0, w.Add(row)
}
