//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Abort — drop the unfinished files
package seg

import "os"

func (w *Writer) Abort() {
	w.f.Close()
	w.tf.Close()
	os.Remove(w.path + ".tmp")
	os.Remove(w.path + ".tbl")
}
