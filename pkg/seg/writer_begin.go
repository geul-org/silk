//ff:func feature=seg type=writer control=sequence
//ff:what Writer.Begin — open the next partition; its key slot must be PartWidth bytes and greater than the previous key
package seg

import (
	"bytes"
	"fmt"
)

// Begin starts partition key (copied).
func (w *Writer) Begin(key []byte) error {
	if w.open || len(key) != w.info.PartWidth || (w.key != nil && bytes.Compare(key, w.key) <= 0) {
		return fmt.Errorf("%w: key %x after %x", ErrOrder, key, w.key)
	}
	w.key, w.open, w.n = append([]byte{}, key...), true, 0
	w.buf = w.buf[:0]
	return nil
}
