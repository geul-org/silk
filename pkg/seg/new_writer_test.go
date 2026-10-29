//ff:func feature=seg type=writer control=sequence
//ff:what TestNewWriter — a writer on path.tmp with counts reset; a directory blocked by a file or a .tmp that is a directory fails
package seg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewWriter(t *testing.T) {
	d := t.TempDir()
	w, err := NewWriter(filepath.Join(d, "x", "a.sseg"), Info{Rows: 9, Parts: 3}, 1)
	if err != nil || w.info.Rows != 0 || w.info.Parts != 0 {
		t.Fatal(err)
	}
	w.Abort()
	os.WriteFile(filepath.Join(d, "file"), nil, 0o644)
	if _, err := NewWriter(filepath.Join(d, "file", "a.sseg"), Info{}, 1); err == nil {
		t.Fatal("parent is a file")
	}
	os.Mkdir(filepath.Join(d, "b.sseg.tmp"), 0o755)
	if _, err := NewWriter(filepath.Join(d, "b.sseg"), Info{}, 1); err == nil {
		t.Fatal(".tmp is a directory")
	}
}
