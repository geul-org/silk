//ff:func feature=recmap type=helper control=sequence
//ff:what failReader.ReadAt — bytes.Reader semantics below the failure offset, an error from it on
package recmap

import (
	"bytes"
	"errors"
)

// ReadAt fails at or after r.from.
func (r failReader) ReadAt(p []byte, off int64) (int, error) {
	if off >= r.from {
		return 0, errors.New("read failure")
	}
	return bytes.NewReader(r.b).ReadAt(p, off)
}
