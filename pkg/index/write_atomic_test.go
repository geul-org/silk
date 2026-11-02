//ff:func feature=index type=scanner control=sequence
//ff:what TestWriteAtomic — files appear whole; a path under a file or a directory in the way fails
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAtomic(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "a", "b")
	if err := writeAtomic(p, []byte("x")); err != nil {
		t.Fatal(err)
	}
	if err := writeAtomic(filepath.Join(p, "c"), nil); err == nil {
		t.Fatal("under a file")
	}
	os.Mkdir(filepath.Join(dir, "d.tmp"), 0o755)
	if err := writeAtomic(filepath.Join(dir, "d"), nil); err == nil {
		t.Fatal("tmp is a directory")
	}
	os.Mkdir(filepath.Join(dir, "e"), 0o755)
	os.WriteFile(filepath.Join(dir, "e", "x"), nil, 0o644)
	if err := writeAtomic(filepath.Join(dir, "e"), nil); err == nil {
		t.Fatal("rename over a full directory")
	}
	b := Manifest{}
	if err := commit(filepath.Join(p, "x"), &b); err == nil {
		t.Fatal("commit under a file")
	}
	os.WriteFile(filepath.Join(dir, currentFile+".tmp"), nil, 0o644)
	os.Mkdir(filepath.Join(dir, currentFile), 0o755)
	os.WriteFile(filepath.Join(dir, currentFile, "x"), nil, 0o644)
	if err := commit(dir, &b); err == nil {
		t.Fatal("CURRENT cannot be replaced")
	}
}
