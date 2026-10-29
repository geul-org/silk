//ff:func feature=recmap type=helper control=sequence
//ff:what offsFail.ReadAt — bytes.Reader semantics, except an error for a read starting at the off[] position
package recmap

import (
	"bytes"
	"errors"
)

// ReadAt fails at r.at.
func (r offsFail) ReadAt(p []byte, off int64) (int, error) {
	if off == r.at {
		return 0, errors.New("off[] read failure")
	}
	return bytes.NewReader(r.b).ReadAt(p, off)
}
