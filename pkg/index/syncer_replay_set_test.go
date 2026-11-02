//ff:func feature=index type=scanner control=sequence
//ff:what TestSyncerReplaySet — a seen pack is skipped; a new one replays and fails with the step
package index

import (
	"bufio"
	"context"
	"path/filepath"
	"testing"
)

func TestSyncerReplaySet(t *testing.T) {
	ctx := context.Background()
	_, dir, a := setup(t, 2)
	u := &updater{ar: a, dir: dir, o: testOpts, m: Manifest{Paths: Section{Name: "paths"}}, tmp: t.TempDir()}
	s := &syncer{u: u, iss: map[string]uint32{}, sp: newSpill(pathInfo(), u.tmp, &u.n)}
	s.dw, _ = newDocWriter(filepath.Join(u.tmp, "d"), 0)
	st := a.Sets[0]
	seen := map[PackRef]bool{{Set: st.Name, Hash: st.Data()[0].Hash}: true}
	if err := s.replaySet(ctx, st, seen); err != nil || s.next != 0 {
		t.Fatal(err, s.next)
	}
	if err := s.replaySet(ctx, st, map[PackRef]bool{}); err != nil || s.next != 2 {
		t.Fatal(err, s.next)
	}
	s.dw.bw = bufio.NewWriterSize(failingWriter{}, 8)
	if err := s.replaySet(ctx, st, map[PackRef]bool{}); err == nil {
		t.Fatal("step failure")
	}
}
