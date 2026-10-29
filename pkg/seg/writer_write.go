//ff:func feature=seg type=writer control=sequence
//ff:what Writer.write — bytes to the file and the running hash
package seg

// write appends b to the segment body.
func (w *Writer) write(b []byte) error {
	if _, err := w.bw.Write(b); err != nil {
		return err
	}
	w.h.Write(b)
	w.written += int64(len(b))
	return nil
}
