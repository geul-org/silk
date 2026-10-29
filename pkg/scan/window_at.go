//ff:func feature=scan type=reader control=sequence
//ff:what window.at — bytes from off with at least n available (fewer only at the pack's end), refilling with one ReadAt when needed
package scan

import "fmt"

// at returns the buffer from off. A refill starts at off and reads
// max(w.min, n) bytes, capped at the pack size.
func (w *window) at(off uint64, n int) ([]byte, error) {
	if off >= w.size {
		return nil, fmt.Errorf("%w: %d beyond pack size %d", ErrRecord, off, w.size)
	}
	need := min(uint64(n), w.size-off)
	if off >= w.lo && off+need <= w.lo+uint64(len(w.buf)) {
		return w.buf[off-w.lo:], nil
	}
	k := min(uint64(max(w.min, n)), w.size-off)
	if uint64(cap(w.buf)) < k {
		w.buf = make([]byte, k)
	}
	w.buf, w.lo = w.buf[:k], off
	if _, err := w.data.ReadAt(w.buf, int64(off)); err != nil {
		w.buf = w.buf[:0]
		return nil, err
	}
	return w.buf, nil
}
