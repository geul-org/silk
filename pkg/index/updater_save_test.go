//ff:func feature=index type=scanner control=sequence
//ff:what TestUpdaterSave — bitmap, commit and collection; each failure surfaces
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUpdaterSave(t *testing.T) {
	_, dir, a := built(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir)}
	g := u.m.Generation
	if err := u.save(); err != nil || u.m.Generation != g+1 {
		t.Fatal(err)
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	bad := &updater{ar: a, dir: file, m: Manifest{}}
	if err := bad.save(); err == nil {
		t.Fatal("under a file")
	}
	full := t.TempDir()
	os.MkdirAll(filepath.Join(full, currentFile, "x"), 0o755)
	cu := &updater{ar: a, dir: full, m: Manifest{}}
	if err := cu.save(); err == nil {
		t.Fatal("CURRENT blocked")
	}
}
