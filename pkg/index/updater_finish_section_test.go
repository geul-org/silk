//ff:func feature=index type=scanner control=sequence
//ff:what TestUpdaterFinishSection — empty rows build nothing; a run is published; a missing one fails
package index

import (
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestUpdaterFinishSection(t *testing.T) {
	_, dir, a := built(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	s := Section{Pending: true}
	var r Report
	if err := u.finishSection(&s, "", &r); err != nil || s.Pending {
		t.Fatal(err)
	}
	p := filepath.Join(u.tmp, "r")
	w, _ := seg.NewWriter(p, pathInfo())
	w.Add(nil, make([]byte, pathKey+4))
	w.Finish()
	s = Section{Pending: true}
	if err := u.finishSection(&s, p, &r); err != nil || r.Rows != 1 || len(s.Runs) != 1 {
		t.Fatal(err)
	}
	if err := u.finishSection(&s, filepath.Join(u.tmp, "none"), &r); err == nil {
		t.Fatal("missing")
	}
}
