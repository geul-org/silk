//ff:func feature=index type=scanner control=sequence
//ff:what TestMergeFiles — merges keep or remove their inputs; missing inputs and unwritable outputs fail
package index

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestMergeFiles(t *testing.T) {
	dir := t.TempDir()
	info := seg.Info{RowWidth: 8, DocOff: 4}
	mk := func(name string, k byte) string {
		p := filepath.Join(dir, name)
		w, _ := seg.NewWriter(p, info)
		w.Add(nil, []byte{0, 0, 0, k, 0, 0, 0, 1})
		w.Finish()
		return p
	}
	a, b := mk("a", 1), mk("b", 2)
	if _, err := mergeKept([]string{a, b}, info, seg.Hooks{}, filepath.Join(dir, "k")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(a); err != nil {
		t.Fatal("kept")
	}
	if _, err := mergeFiles([]string{a, b}, info, seg.Hooks{}, filepath.Join(dir, "m")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(a); err == nil {
		t.Fatal("removed")
	}
	if _, err := mergeFiles([]string{a}, info, seg.Hooks{}, filepath.Join(dir, "x")); err == nil {
		t.Fatal("missing input")
	}
	if _, err := mergeKept([]string{a}, info, seg.Hooks{}, filepath.Join(dir, "x")); err == nil {
		t.Fatal("missing input")
	}
	c := mk("c", 3)
	if _, err := mergeFiles([]string{c}, info, seg.Hooks{}, filepath.Join(c, "x")); err == nil {
		t.Fatal("bad output")
	}
	if _, err := mergeKept([]string{c}, info, seg.Hooks{}, filepath.Join(c, "x")); err == nil {
		t.Fatal("bad output")
	}
	other := seg.Info{RowWidth: 8, DocOff: 4, PartWidth: 2}
	if _, err := mergeFiles([]string{c}, other, seg.Hooks{}, filepath.Join(dir, "y")); err == nil {
		t.Fatal("shape")
	}
	if _, err := mergeKept([]string{c}, other, seg.Hooks{}, filepath.Join(dir, "z")); err == nil {
		t.Fatal("shape")
	}
	if p, err := mergeDown(nil, info, 2, seg.Hooks{}, nil); p != "" || err != nil {
		t.Fatal(p, err)
	}
	boom := filepath.Join(dir, "none")
	if _, err := mergeDown([]string{boom, boom}, info, 2, seg.Hooks{}, func() string { return filepath.Join(dir, "w") }); err == nil {
		t.Fatal("missing")
	}
	if _, err := mergeLevel([]string{boom}, info, 2, seg.Hooks{}, func() string { return filepath.Join(dir, "w") }); err == nil {
		t.Fatal("missing")
	}
	if _, err := openSegs([]string{c, boom}); err == nil {
		t.Fatal("missing")
	}
	if _, err := infoOf(boom); err == nil {
		t.Fatal("missing")
	}
	removeAll([]string{boom})
}
