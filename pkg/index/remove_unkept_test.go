//ff:func feature=index type=writer control=sequence
//ff:what TestRemoveUnkept — kept names stay, others are removed and counted; an entry that cannot be removed is not counted
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveUnkept(t *testing.T) {
	d := t.TempDir()
	sd := filepath.Join(d, "seg")
	os.MkdirAll(filepath.Join(sd, "full", "x"), 0o755)
	os.WriteFile(filepath.Join(sd, "a"), nil, 0o644)
	os.WriteFile(filepath.Join(sd, "b"), nil, 0o644)
	es, _ := os.ReadDir(sd)
	if n := removeUnkept(d, "seg", es, map[string]bool{"seg/a": true}); n != 1 {
		t.Fatal(n)
	}
	if _, err := os.Stat(filepath.Join(sd, "a")); err != nil {
		t.Fatal(err)
	}
}
