//ff:func feature=query type=helper control=sequence
//ff:what tailFail.ReadAt — bytes below the failure offset, an error from it on
package query

import (
	"bytes"
	"errors"
)

// ReadAt fails at or after t.from.
func (t tailFail) ReadAt(p []byte, off int64) (int, error) {
	if off >= t.from {
		return 0, errors.New("read failure")
	}
	return bytes.NewReader(t.b).ReadAt(p, off)
}
