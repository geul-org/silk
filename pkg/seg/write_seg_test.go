//ff:func feature=seg type=helper control=iteration dimension=1
//ff:what writeSeg — a segment of partitions pa (rows 1,2 docs 1,2) and pb (row 3 doc 3)
package seg

import (
	"path/filepath"
	"testing"
)

func writeSeg(t *testing.T, name string) *File {
	t.Helper()
	p := filepath.Join(t.TempDir(), name)
	w, err := NewWriter(p, testInfo())
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range []struct {
		k string
		r []byte
	}{{"pa", row(0, 1, 1)}, {"pa", row(0, 2, 2)}, {"pb", row(0, 3, 3)}} {
		if err := w.Add([]byte(e.k), e.r); err != nil {
			t.Fatal(err)
		}
	}
	if n, err := w.Finish(); err != nil || n != 3 {
		t.Fatal(n, err)
	}
	f, err := Open(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { f.Close() })
	return f
}
