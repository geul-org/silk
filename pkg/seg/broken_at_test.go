//ff:func feature=seg type=helper control=sequence
//ff:what brokenAt — test helper: a broken Writer (Min 0) with partition "aa" open, so the next row fails to write
package seg

import "testing"

// brokenAt prepares a Writer whose next Add fails.
func brokenAt(t testing.TB) *Writer {
	w := brokenWriter(t, 0)
	w.Begin([]byte("aa"))
	w.write(make([]byte, 15))
	return w
}
