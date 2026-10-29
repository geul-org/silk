//ff:func feature=scan type=helper control=sequence
//ff:what failAfter.ReadAt — pass through while reads remain, then an error
package scan

import "errors"

// ReadAt counts down.
func (f *failAfter) ReadAt(p []byte, off int64) (int, error) {
	if f.n <= 0 {
		return 0, errors.New("read failure")
	}
	f.n--
	return f.r.ReadAt(p, off)
}
