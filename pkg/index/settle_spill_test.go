//ff:func feature=index type=scanner control=sequence
//ff:what TestSettleSpill — a run that cannot be made stops settling
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSettleSpill(t *testing.T) {
	_, dir, a := built(t, 1)
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: file}
	s := &syncer{u: u, sp: newSpill(pathInfo(), file, &u.n)}
	s.sp.buf.Add(nil, make([]byte, pathKey+4))
	if err := s.settle(); err == nil {
		t.Fatal("under a file")
	}
	u = &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	s = &syncer{u: u, sp: newSpill(pathInfo(), u.tmp, &u.n)}
	s.sp.buf.Add(nil, make([]byte, pathKey+4))
	s.sp.buf.Add(nil, make([]byte, pathKey+4))
	if err := s.settle(); err != nil {
		t.Fatal(err)
	}
	s = &syncer{u: u, sp: newSpill(pathInfo(), u.tmp, &u.n)}
	s.sp.runs = []string{filepath.Join(u.tmp, "none")}
	if err := s.settle(); err == nil {
		t.Fatal("unreadable run")
	}
}
