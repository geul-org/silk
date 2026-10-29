//ff:func feature=cli type=command control=sequence
//ff:what TestIndexBuild — builds and reports; a bad spec file, a non-archive and a locked index fail
package cli

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestIndexBuild(t *testing.T) {
	data := fixturePacked(t)
	dir := filepath.Join(data, ".silk")
	if code := indexBuild(data, dir, indexFlags{}, io.Discard, io.Discard); code != 0 {
		t.Fatal(code)
	}
	if code := indexBuild(data, dir, indexFlags{spec: "/nonexistent"}, io.Discard, io.Discard); code != 1 {
		t.Fatal("spec", code)
	}
	if code := indexBuild(t.TempDir(), dir, indexFlags{}, io.Discard, io.Discard); code != 1 {
		t.Fatal("archive", code)
	}
	os.WriteFile(filepath.Join(dir, "LOCK"), []byte("1\n"), 0o644)
	if code := indexBuild(data, dir, indexFlags{}, io.Discard, io.Discard); code != 1 {
		t.Fatal("locked", code)
	}
}
