//ff:func feature=index type=scanner control=sequence
//ff:what TestCompactorPlace — a run moves into seg/; a missing run or a blocked directory fails
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCompactorPlace(t *testing.T) {
	dir := t.TempDir()
	os.Mkdir(filepath.Join(dir, dirSeg), 0o755)
	p := filepath.Join(dir, "r")
	os.WriteFile(p, []byte("x"), 0o644)
	c := &compactor{dir: dir}
	if r, err := c.place(p, 1); err != nil || r.Bytes != 1 {
		t.Fatal(r, err)
	}
	if _, err := c.place(p, 1); err == nil {
		t.Fatal("gone")
	}
}
