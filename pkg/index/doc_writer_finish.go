//ff:func feature=index type=writer control=sequence
//ff:what docWriter.finish — tail, sync, rename
package index

import (
	"encoding/binary"
	"errors"
	"os"
)

func (w *docWriter) finish() error {
	t := append([]byte(docMagic), binary.BigEndian.AppendUint32(binary.BigEndian.AppendUint32(nil, w.first), w.count)...)
	t = append(append(t, make([]byte, 8)...), w.sum.Sum(nil)...)
	_, err := w.bw.Write(t)
	if err := errors.Join(err, w.bw.Flush(), w.f.Sync(), w.f.Close()); err != nil {
		return err
	}
	return os.Rename(w.path+".tmp", w.path)
}
