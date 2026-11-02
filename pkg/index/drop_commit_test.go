//ff:func feature=index type=scanner control=sequence
//ff:what TestDropCommit — a commit that fails is reported
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDropCommit(t *testing.T) {
	_, dir, _ := setup(t, 1)
	os.Remove(filepath.Join(dir, currentFile))
	os.MkdirAll(filepath.Join(dir, currentFile, "x"), 0o755)
	if err := Drop(dir, "s"); err == nil {
		t.Fatal("no current")
	}
	d := t.TempDir()
	m := Manifest{Sections: []Section{{Name: "s"}}}
	commit(d, &m)
	os.Remove(filepath.Join(d, manifestName(m.Generation+1)))
	os.MkdirAll(filepath.Join(d, manifestName(m.Generation+1), "x"), 0o755)
	if err := Drop(d, "s"); err == nil {
		t.Fatal("commit blocked")
	}
}
