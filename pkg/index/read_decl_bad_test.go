//ff:func feature=index type=scanner control=sequence
//ff:what TestReadDeclBad — damaged JSON is ErrFormat
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadDeclBad(t *testing.T) {
	dir := t.TempDir()
	os.Mkdir(filepath.Join(dir, dirDecl), 0o755)
	os.WriteFile(filepath.Join(dir, dirDecl, "x.json"), []byte("["), 0o644)
	if _, err := readDecl(dir, "x"); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(dir, dirDecl, "y.json"), []byte("{}"), 0o644)
	if _, err := readDecl(dir, "y"); err != nil {
		t.Fatal(err)
	}
}
