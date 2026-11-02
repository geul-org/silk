//ff:func feature=index type=scanner control=sequence
//ff:what TestSyncerReplayErrors — seen packs are skipped, and a set's error stops the walk
package index

import (
	"bufio"
	"context"
	"path/filepath"
	"testing"
)

func TestSyncerReplayErrors(t *testing.T) {
	ctx := context.Background()
	_, dir, a := setup(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: Manifest{Paths: Section{Name: "paths"}}, tmp: t.TempDir()}
	s := &syncer{u: u, iss: map[string]uint32{}, sp: newSpill(pathInfo(), u.tmp, &u.n)}
	s.dw, _ = newDocWriter(filepath.Join(u.tmp, "d"), 0)
	s.dw.bw = bufio.NewWriterSize(failingWriter{}, 1)
	if err := s.replay(ctx); err == nil {
		t.Fatal("step error")
	}
	s.u.m.Packs = []PackRef{{Set: a.Sets[0].Name, Hash: a.Sets[0].Data()[0].Hash}}
	if err := s.replay(ctx); err != nil {
		t.Fatal(err)
	}
}
