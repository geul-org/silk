//ff:func feature=index type=builder control=sequence
//ff:what TestUpdaterWithAppends — a run merges with appends; a missing run or a blocked output fails
package index

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestUpdaterWithAppends(t *testing.T) {
	u := &updater{dir: t.TempDir(), tmp: t.TempDir()}
	u.live.Put(0, true)
	p := filepath.Join(u.tmp, "r")
	w, _ := seg.NewWriter(p, pathInfo())
	w.Add(nil, make([]byte, pathKey+4))
	w.Finish()
	if _, n, err := u.withAppends(&Section{}, p); n != 1 || err != nil {
		t.Fatal(n, err)
	}
	if _, _, err := u.withAppends(&Section{Appends: []SegRef{{File: "none"}}}, p); err == nil {
		t.Fatal("missing append")
	}
	if _, _, err := u.withAppends(&Section{}, p+"x"); err == nil {
		t.Fatal("missing run")
	}
	u.tmp = filepath.Join(p, "x")
	if _, _, err := u.withAppends(&Section{}, p); err == nil {
		t.Fatal("output under a file")
	}
}
