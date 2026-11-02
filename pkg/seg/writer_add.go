//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Add — append a row; partitions and the rows inside each must ascend
package seg

import (
	"bytes"
	"fmt"
)

func (w *Writer) Add(part, row []byte) error {
	if len(part) != w.info.PartWidth || len(row) != w.info.RowWidth {
		return fmt.Errorf("%w: partition %d, row %d bytes", ErrFormat, len(part), len(row))
	}
	c := bytes.Compare(part, w.part)
	if w.open && (c < 0 || c == 0 && bytes.Compare(row, w.last) < 0) {
		return ErrOrder
	}
	if !w.open || c != 0 {
		if err := w.nextPart(part); err != nil {
			return err
		}
	}
	copy(w.last, row)
	w.rows++
	w.sum.Write(row)
	_, err := w.bw.Write(row)
	return err
}
