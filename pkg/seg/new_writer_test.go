//ff:func feature=seg type=scanner control=sequence
//ff:what TestNewWriter — bad shapes and unusable paths fail
package seg

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestNewWriter(t *testing.T) {
	dir := t.TempDir()
	if _, err := NewWriter(filepath.Join(dir, "a"), Info{}); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	file := filepath.Join(dir, "file")
	os.WriteFile(file, nil, 0o644)
	if _, err := NewWriter(filepath.Join(file, "x", "a"), testInfo()); err == nil {
		t.Fatal("under a file")
	}
	os.Mkdir(filepath.Join(dir, "d.tmp"), 0o755)
	if _, err := NewWriter(filepath.Join(dir, "d"), testInfo()); err == nil {
		t.Fatal("tmp is a directory")
	}
	os.Mkdir(filepath.Join(dir, "e.tbl"), 0o755)
	if _, err := NewWriter(filepath.Join(dir, "e"), testInfo()); err == nil {
		t.Fatal("tbl is a directory")
	}
}
