//ff:func feature=index type=scanner control=sequence
//ff:what TestDocWriterFinish — a rename blocked by a directory and a closed file fail
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDocWriterFinish(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "b", "x"), 0o755)
	w, _ := newDocWriter(filepath.Join(dir, "b"), 0)
	if err := w.finish(); err == nil {
		t.Fatal("rename over a directory")
	}
	w, _ = newDocWriter(filepath.Join(dir, "c"), 0)
	w.f.Close()
	if err := w.finish(); err == nil {
		t.Fatal("closed")
	}
}
