//ff:func feature=seg type=writer control=sequence
//ff:what Writer.nextPart — close the open partition's table entry and start part
package seg

func (w *Writer) nextPart(part []byte) error {
	if err := w.closePart(); err != nil {
		return err
	}
	w.part = append(w.part[:0], part...)
	w.open = true
	return nil
}
