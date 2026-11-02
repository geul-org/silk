//ff:func feature=seg type=writer control=sequence
//ff:what Writer.copyTable — append the spilled partition table after the rows
package seg

import "io"

func (w *Writer) copyTable() error {
	if _, err := w.tf.Seek(0, io.SeekStart); err != nil {
		return err
	}
	_, err := io.Copy(io.MultiWriter(w.bw, w.sum), w.tf)
	return err
}
