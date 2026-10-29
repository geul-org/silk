//ff:func feature=seg type=reader control=sequence
//ff:what failReader.ReadAt — read through until the budget is spent, then io.ErrUnexpectedEOF
package seg

import "io"

// ReadAt fails once n reads have passed.
func (f *failReader) ReadAt(b []byte, off int64) (int, error) {
	if f.n <= 0 {
		return 0, io.ErrUnexpectedEOF
	}
	f.n--
	return f.r.ReadAt(b, off)
}
