//ff:func feature=index type=scanner control=sequence
//ff:what TestReadDecl — declarations and grammars load; damaged or missing ones fail
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadDecl(t *testing.T) {
	dir := t.TempDir()
	if _, err := readDecl(dir, "x"); err == nil {
		t.Fatal("missing")
	}
	os.MkdirAll(filepath.Join(dir, dirDecl), 0o755)
	os.WriteFile(filepath.Join(dir, dirDecl, "x.json"), []byte("{"), 0o644)
	if _, err := readDecl(dir, "x"); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if g, err := loadGrammar(dir, ""); err != nil || g == nil {
		t.Fatal(err)
	}
	if _, err := loadGrammar(dir, "x"); err == nil {
		t.Fatal("missing")
	}
	os.MkdirAll(filepath.Join(dir, dirGrammar), 0o755)
	os.WriteFile(filepath.Join(dir, dirGrammar, "x.json"), []byte("{"), 0o644)
	if _, err := loadGrammar(dir, "x"); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
