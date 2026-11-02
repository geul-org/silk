//ff:func feature=seg type=scanner control=sequence
//ff:what TestMerge — inputs merge in order, dead rows drop, empty inputs and pauses work
package seg

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestMerge(t *testing.T) {
	a, b := writeSeg(t, "a"), writeSeg(t, "b")
	e := filepath.Join(t.TempDir(), "e")
	w, _ := NewWriter(e, testInfo())
	w.Finish()
	empty, _ := Open(e)
	defer empty.Close()
	p := filepath.Join(t.TempDir(), "m")
	w, _ = NewWriter(p, testInfo())
	pauses := 0
	hk := Hooks{Live: func(d uint32) bool { return d != 2 }, Pause: func() error { pauses++; return nil }}
	n, err := Merge(w, []*File{a, b, empty}, hk)
	if err != nil || n != 2 || pauses != 1 {
		t.Fatal(n, err, pauses)
	}
	w.Finish()
	m, _ := Open(p)
	defer m.Close()
	if m.Rows != 4 || m.Parts != 2 || m.Verify() != nil {
		t.Fatal(m.Rows)
	}
	w, _ = NewWriter(filepath.Join(t.TempDir(), "n"), Info{PartWidth: 1, RowWidth: 6, DocOff: 2})
	defer w.Abort()
	if _, err := Merge(w, []*File{a}, Hooks{}); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	boom := errors.New("stop")
	if _, err := Merge(w, []*File{a}, Hooks{Pause: func() error { return boom }}); !errors.Is(err, boom) {
		t.Fatal(err)
	}
}
