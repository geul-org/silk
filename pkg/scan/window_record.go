//ff:func feature=scan type=reader control=sequence
//ff:what window.record — exactly the record at a pack offset: read its TAG WORD (and Length extension), then its whole length
package scan

import "fmt"

// record frames the record at off with RecLen.
func (w *window) record(off uint64) ([]byte, error) {
	b, err := w.at(off, maxHead)
	if err != nil {
		return nil, err
	}
	n := RecLen(b)
	if n <= 0 {
		return nil, fmt.Errorf("%w: %d", ErrRecord, off)
	}
	if b, err = w.at(off, n); err != nil {
		return nil, err
	}
	if len(b) < n {
		return nil, fmt.Errorf("%w: %d runs past the pack", ErrRecord, off)
	}
	return b[:n], nil
}
