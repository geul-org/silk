//ff:func feature=archive type=helper control=sequence
//ff:what TestWellKnown — a data folder resolves to its .well-known/geul; a directory without one (or with a file there) is itself
package archive

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWellKnown(t *testing.T) {
	dir := t.TempDir()
	if got := wellKnown(dir); got != dir {
		t.Fatal(got)
	}
	os.MkdirAll(filepath.Join(dir, ".well-known"), 0o755)
	os.WriteFile(filepath.Join(dir, ".well-known", "geul"), nil, 0o644)
	if got := wellKnown(dir); got != dir {
		t.Fatal(got)
	}
	dir2 := t.TempDir()
	wk := filepath.Join(dir2, ".well-known", "geul")
	os.MkdirAll(wk, 0o755)
	if got := wellKnown(dir2); got != wk {
		t.Fatal(got)
	}
}
