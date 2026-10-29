//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Finish — write the other rows, the partition table and the footer, sync, and rename the file into place; its Info
package seg

import (
	"errors"
	"os"
)

// Finish completes the segment. The footer's Sum covers every byte before
// it. On error the temporary file is removed.
func (w *Writer) Finish() (Info, error) {
	tail := appendTable(w.other, w.parts)
	w.h.Write(tail)
	w.info.Parts = uint32(len(w.parts))
	copy(w.info.Sum[:], w.h.Sum(nil))
	_, werr := w.bw.Write(append(tail, encodeFooter(w.info)...))
	w.written += int64(len(tail) + footerSize)
	if err := errors.Join(werr, w.bw.Flush(), w.f.Sync(), w.f.Close()); err != nil {
		os.Remove(w.Path + ".tmp")
		return Info{}, err
	}
	w.info.Bytes = w.written
	return w.info, os.Rename(w.Path+".tmp", w.Path)
}
