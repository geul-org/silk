//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestSyncer — issuers number once; a failing writer stops a step; a pack whose inventory cannot be read stops the replay
package index

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

func TestSyncer(t *testing.T) {
	ctx := context.Background()
	_, dir, a := setup(t, 2)
	u := &updater{ar: a, dir: dir, o: testOpts, m: Manifest{Paths: Section{Name: "paths"}}, tmp: t.TempDir()}
	s := &syncer{u: u, iss: map[string]uint32{}}
	if s.issuer("a") != 0 || s.issuer("b") != 1 || s.issuer("a") != 0 || len(u.m.Issuers) != 2 {
		t.Fatal(u.m.Issuers)
	}
	if len(identity(DocEntry{Issuer: 1, ID: 2, Version: 3})) != pathKey {
		t.Fatal("identity")
	}
	s.dw, _ = newDocWriter(filepath.Join(u.tmp, "d"), 0)
	s.sp = newSpill(pathInfo(), filepath.Join(u.tmp, "s"), &u.n)
	u.o.Limit = 1
	if err := s.step(archive.Step{Issuer: "a", ID: 1, Version: 1}); err != nil {
		t.Fatal(err)
	}
	if err := s.step(archive.Step{Issuer: "a", ID: 1, Version: 1, Deleted: true}); err != nil {
		t.Fatal(err)
	}
	if u.live.Get(1) || !u.live.Get(0) || len(s.sp.runs) != 2 {
		t.Fatal(u.live, s.sp.runs)
	}
	s.dw.f.Close()
	s.dw.bw = bufio.NewWriterSize(failingWriter{}, 8)
	if err := s.step(archive.Step{Issuer: "a"}); err == nil {
		t.Fatal("closed table")
	}
	s = &syncer{u: u, iss: map[string]uint32{}, sp: newSpill(pathInfo(), u.tmp, &u.n)}
	boom := filepath.Join(a.Sets[0].Dir, "o")
	os.Chmod(boom, 0o755)
	es, _ := os.ReadDir(boom)
	for _, e := range es {
		os.Rename(filepath.Join(boom, e.Name()), filepath.Join(boom, e.Name()+".away"))
	}
	if err := s.replay(ctx); err == nil {
		t.Fatal("no pack file")
	}
	if err := (&syncer{u: &updater{dir: u.tmp, m: Manifest{}}}).settleEmpty(); err != nil {
		t.Fatal(err)
	}
}
