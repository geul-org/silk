//ff:func feature=index type=scanner control=sequence
//ff:what TestWriteAtomicClosed — a target blocked by a full directory fails at rename
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAtomicClosed(t *testing.T) {
	dir := t.TempDir()
	if err := writeAtomic(filepath.Join(dir, "a"), []byte("x")); err != nil {
		t.Fatal(err)
	}
	os.Symlink("/dev/full", filepath.Join(dir, "full.tmp"))
	if err := writeAtomic(filepath.Join(dir, "full"), []byte("x")); err == nil {
		t.Fatal("write to a full disk")
	}
	os.MkdirAll(filepath.Join(dir, "b", "x"), 0o755)
	if err := writeAtomic(filepath.Join(dir, "b"), []byte("x")); err == nil {
		t.Fatal("rename")
	}
}
