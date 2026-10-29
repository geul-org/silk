//ff:func feature=seg type=writer control=iteration dimension=1
//ff:what TestWritePart — rows sorted for a run, kept in arrival order for an appended segment; a bad key fails
package seg

import (
	"path/filepath"
	"testing"
)

func TestWritePart(t *testing.T) {
	rows := append(row("k2", 1, 0), row("k1", 1, 0)...)
	for _, sorted := range []bool{true, false} {
		p := filepath.Join(t.TempDir(), "w.sseg")
		w, _ := NewWriter(p, Info{PartWidth: 2, SortWidth: 2}, 0)
		if err := writePart(w, []byte("aa"), rows, sorted); err != nil {
			t.Fatal(err)
		}
		w.Finish()
		f, _ := Open(p)
		b, _ := f.readRows(f.partRegion(0), 0, 1)
		f.Close()
		if want := map[bool]string{true: "k1", false: "k2"}[sorted]; string(b[:2]) != want {
			t.Fatal(sorted, string(b[:2]))
		}
	}
	w, _ := NewWriter(filepath.Join(t.TempDir(), "w.sseg"), Info{PartWidth: 2, SortWidth: 2}, 0)
	defer w.Abort()
	if err := writePart(w, []byte("a"), rows, true); err == nil {
		t.Fatal("bad key")
	}
}
