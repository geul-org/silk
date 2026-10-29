//ff:func feature=seg type=helper control=sequence
//ff:what Writer.Written — bytes written so far (after Finish, the file size)
package seg

// Written counts the bytes handed to the file.
func (w *Writer) Written() int64 {
	return w.written
}
