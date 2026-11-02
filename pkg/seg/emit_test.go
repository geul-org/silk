//ff:func feature=seg type=scanner control=sequence
//ff:what TestEmit — a dead row is skipped, a live one written
package seg

import (
	"path/filepath"
	"testing"
)

func TestEmit(t *testing.T) {
	a := writeSeg(t, "a")
	c := &cursor{f: a}
	c.load()
	w, _ := NewWriter(filepath.Join(t.TempDir(), "x"), testInfo())
	defer w.Abort()
	if keep, err := emit(w, c, Hooks{Live: func(uint32) bool { return false }}); keep || err != nil {
		t.Fatal(keep, err)
	}
	if keep, err := emit(w, c, Hooks{}); !keep || err != nil || w.rows != 1 {
		t.Fatal(keep, err)
	}
}
