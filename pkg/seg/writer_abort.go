//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Abort — drop an unfinished segment (close and remove its temporary file)
package seg

import "os"

// Abort discards the temporary file.
func (w *Writer) Abort() {
	w.f.Close()
	os.Remove(w.Path + ".tmp")
}
