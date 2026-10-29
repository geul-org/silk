//ff:func feature=seg type=writer control=sequence
//ff:what TestWriterAdd — rows stay buffered below Min, then the buffer and later rows are written; write errors surface
package seg

import (
	"path/filepath"
	"testing"
)

func TestWriterAdd(t *testing.T) {
	w := brokenWriter(t, 3)
	w.Begin([]byte("aa"))
	if err := w.Add(row("k1", 1, 0)); err != nil || len(w.buf) != 10 {
		t.Fatal(err, len(w.buf))
	}
	w.Add(row("k2", 1, 2))
	if err := w.Add(row("k3", 1, 4)); err == nil {
		t.Fatal("flushing the buffer to a closed file")
	}
	ok, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{PartWidth: 2, SortWidth: 2}, 2)
	defer ok.Abort()
	ok.Begin([]byte("aa"))
	ok.Add(row("k1", 1, 0))
	if err := ok.Add(row("k2", 1, 0)); err != nil || len(ok.buf) != 0 || ok.written != 20 {
		t.Fatal(err, ok.written)
	}
	w2 := brokenWriter(t, 0)
	w2.Begin([]byte("aa"))
	w2.Add(row("k1", 1, 0))
	w2.Add(row("k2", 1, 0))
	if err := w2.Add(row("k3", 1, 0)); err == nil {
		t.Fatal("direct write to a closed file")
	}
}
