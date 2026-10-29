//ff:func feature=seg type=writer control=sequence
//ff:what TestEmitRow — a dead row counts 1, a live one is added, a write error is returned
package seg

import (
	"path/filepath"
	"testing"
)

func TestEmitRow(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "e.sseg"), Info{PartWidth: 2, SortWidth: 2}, 0)
	defer w.Abort()
	w.Begin([]byte("aa"))
	if d, err := emitRow(w, row("k1", 1, 0), func(uint32) bool { return false }); d != 1 || err != nil || w.n != 0 {
		t.Fatal(d, err)
	}
	if d, err := emitRow(w, row("k1", 1, 0), allLive); d != 0 || err != nil || w.n != 1 {
		t.Fatal(d, err)
	}
	if _, err := emitRow(brokenAt(t), row("k1", 1, 0), allLive); err == nil {
		t.Fatal("write error")
	}
}
