//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestSettle — an unreadable new run or old segment stops settling
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSettle(t *testing.T) {
	_, dir, a := built(t, 2)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	s := &syncer{u: u, sp: newSpill(pathInfo(), u.tmp, &u.n)}
	s.sp.buf.Add(nil, make([]byte, pathKey+4))
	for _, r := range append(u.m.Paths.Runs, u.m.Paths.Appends...) {
		os.Remove(filepath.Join(dir, r.File))
	}
	if err := s.settle(); err == nil {
		t.Fatal("old segment gone")
	}
	if p := segPaths("d", Section{Runs: []SegRef{{File: "a"}}, Appends: []SegRef{{File: "b"}}}); len(p) != 2 {
		t.Fatal(p)
	}
}
