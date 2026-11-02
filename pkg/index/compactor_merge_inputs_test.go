//ff:func feature=index type=scanner control=sequence
//ff:what TestCompactorMergeInputs — several levels under a small fan-in; failures at each level
package index

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestCompactorMergeInputs(t *testing.T) {
	_, dir, _ := built(t, 2)
	m := current(t, dir)
	c := &compactor{dir: dir, o: Options{Limit: 1}, m: m, tmp: t.TempDir()}
	p := filepath.Join(dir, m.Sections[1].Runs[0].File)
	info, _ := infoOf(p)
	q, rows, err := c.mergeInputs([]string{p, p, p, p, p}, info, seg.Hooks{})
	if err != nil || rows != 10 || q == "" {
		t.Fatal(rows, err)
	}
	if _, _, err := c.mergeInputs([]string{p + "x"}, info, seg.Hooks{}); err == nil {
		t.Fatal("missing input")
	}
	boom := errors.New("stop")
	n := 0
	hk := seg.Hooks{Pause: func() error {
		n++
		if n > 2 {
			return boom
		}
		return nil
	}}
	if _, _, err := c.mergeInputs([]string{p, p, p, p, p}, info, hk); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	n = 0
	hk.Pause = func() error {
		n++
		if n > 3 {
			return boom
		}
		return nil
	}
	if _, _, err := c.mergeInputs([]string{p, p, p, p, p}, info, hk); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if _, _, err := (&compactor{o: Options{Limit: 1}, tmp: filepath.Join(p, "x")}).mergeInputs([]string{p}, info, seg.Hooks{}); err == nil {
		t.Fatal("tmp under a file")
	}
}
