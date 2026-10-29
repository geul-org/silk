//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Suspend — checkpoint the Writer and close its file, keeping the temporary file for ResumeWriter
package seg

import "errors"

// Suspend ends this process's use of the Writer.
func (w *Writer) Suspend() (State, error) {
	st, err := w.Checkpoint()
	return st, errors.Join(err, w.f.Close())
}
