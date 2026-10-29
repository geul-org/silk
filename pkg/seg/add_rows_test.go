//ff:func feature=seg type=writer control=sequence
//ff:what TestAddRows — rows of the writer's width added, dead ones counted; a write error stops
package seg

import (
	"path/filepath"
	"testing"
)

func TestAddRows(t *testing.T) {
	w, _ := NewWriter(filepath.Join(t.TempDir(), "a.sseg"), Info{PartWidth: 2, SortWidth: 2}, 0)
	defer w.Abort()
	w.Begin([]byte("aa"))
	b := append(row("k1", 1, 0), row("k2", 2, 0)...)
	if d, err := addRows(w, b, func(doc uint32) bool { return doc == 1 }); d != 1 || err != nil || w.n != 1 {
		t.Fatal(d, err)
	}
	if _, err := addRows(brokenAt(t), b, allLive); err == nil {
		t.Fatal("write error")
	}
}
