//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestMergeDownLevels — several levels end in one; a failure stops
package index

import (
	"fmt"
	"path/filepath"
	"slices"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestMergeDownLevels(t *testing.T) {
	dir := t.TempDir()
	info := seg.Info{RowWidth: 8, DocOff: 4}
	n := 0
	name := func() string { n++; return filepath.Join(dir, fmt.Sprintf("m%d", n)) }
	var ps []string
	for i := range 5 {
		p := filepath.Join(dir, fmt.Sprintf("r%d", i))
		w, _ := seg.NewWriter(p, info)
		w.Add(nil, []byte{0, 0, 0, byte(i), 0, 0, 0, 1})
		w.Finish()
		ps = append(ps, p)
	}
	p, err := mergeDown(slices.Clone(ps[:4]), info, 2, seg.Hooks{}, name)
	if err != nil || n != 3 {
		t.Fatal(p, n, err)
	}
	next, err := mergeLevel([]string{ps[4]}, info, 2, seg.Hooks{}, name)
	if err != nil || len(next) != 1 {
		t.Fatal(next, err)
	}
	if _, err := mergeDown([]string{next[0], next[0] + "x"}, info, 2, seg.Hooks{}, name); err == nil {
		t.Fatal("missing")
	}
	if _, err := mergeLevel([]string{"none"}, info, 2, seg.Hooks{}, name); err == nil {
		t.Fatal("missing")
	}
}
