//ff:func feature=index type=scanner control=sequence
//ff:what TestDocWriterAdd — entries count and write; a failing buffer surfaces at finish
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDocWriterAdd(t *testing.T) {
	dir := t.TempDir()
	w, err := newDocWriter(filepath.Join(dir, "a"), 5)
	if err != nil {
		t.Fatal(err)
	}
	if err := w.add(DocEntry{ID: 1}); err != nil || w.count != 1 {
		t.Fatal(err)
	}
	if err := w.finish(); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(dir, "a"))
	first, count, body, err := readTail("a", b, docMagic)
	if err != nil || first != 5 || count != 1 || len(body) != docSize {
		t.Fatal(first, count, err)
	}
}
