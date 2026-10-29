//ff:func feature=seg type=helper control=sequence
//ff:what brokenWriter — test helper: a Writer whose file is closed and whose buffer flushes every 16 bytes, so writes fail
package seg

import (
	"bufio"
	"path/filepath"
	"testing"
)

// brokenWriter has PartWidth and SortWidth 2 and Min min.
func brokenWriter(t testing.TB, min int) *Writer {
	t.Helper()
	w, err := NewWriter(filepath.Join(t.TempDir(), "b.sseg"), Info{PartWidth: 2, SortWidth: 2}, min)
	if err != nil {
		t.Fatal(err)
	}
	w.f.Close()
	w.bw = bufio.NewWriterSize(w.f, 16)
	return w
}
