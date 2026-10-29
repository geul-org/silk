//ff:func feature=index type=writer control=iteration dimension=1
//ff:what TestRemoveStale — unlisted maps go, listed maps and other files stay; a failing removal surfaces
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveStale(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "a"), 0o755)
	for _, n := range []string{"keep.srm", "drop.srm", "note.txt"} {
		os.WriteFile(filepath.Join(dir, "a", n), nil, 0o644)
	}
	if err := removeStale(dir, []string{"a/keep.srm"}); err != nil {
		t.Fatal(err)
	}
	for n, want := range map[string]bool{"keep.srm": true, "drop.srm": false, "note.txt": true} {
		if _, err := os.Stat(filepath.Join(dir, "a", n)); (err == nil) != want {
			t.Fatal(n, err)
		}
	}
	os.MkdirAll(filepath.Join(dir, "a", "busy.srm", "x"), 0o755)
	if err := removeStale(dir, nil); err == nil {
		t.Fatal("non-empty directory named like a map")
	}
}
