//ff:func feature=seg type=writer control=sequence
//ff:what Writer.closePart — write the table entry of the open partition, if any
package seg

import "encoding/binary"

func (w *Writer) closePart() error {
	if !w.open {
		return nil
	}
	e := binary.BigEndian.AppendUint64(append([]byte(nil), w.part...), w.start)
	e = binary.BigEndian.AppendUint64(e, w.rows-w.start)
	w.start = w.rows
	w.parts++
	w.open = false
	_, err := w.tw.Write(e)
	return err
}
