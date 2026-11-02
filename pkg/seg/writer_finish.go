//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Finish — table, footer, sync, rename; the row count
package seg

import (
	"errors"
	"os"
)

func (w *Writer) Finish() (uint64, error) {
	if err := errors.Join(w.closePart(), w.tw.Flush()); err != nil {
		w.Abort()
		return 0, err
	}
	if err := w.copyTable(); err != nil {
		w.Abort()
		return 0, err
	}
	_, werr := w.bw.Write(encodeFooter(w.info, w.parts, w.rows, w.sum.Sum(nil)))
	if err := errors.Join(werr, w.bw.Flush(), w.f.Sync(), w.f.Close(), w.tf.Close(), os.Remove(w.path+".tbl")); err != nil {
		os.Remove(w.path + ".tmp")
		return 0, err
	}
	return w.rows, os.Rename(w.path+".tmp", w.path)
}
