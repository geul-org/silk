//ff:func feature=index type=scanner control=sequence
//ff:what TestIngesterParts — feed, doc and sink behaviour: a frame error, the first error kept, parallel errors
package index

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestIngesterParts(t *testing.T) {
	boom := errors.New("boom")
	if err := parallel(4, 2, func(i int) error {
		if i == 1 {
			return boom
		}
		return nil
	}); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if err := parallel(2, 1, func(int) error { return nil }); err != nil {
		t.Fatal(err)
	}
	_, dir, a := built(t, 2)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	g := &ingester{u: u}
	g.fail(boom)
	g.fail(errors.New("later"))
	if !errors.Is(g.err, boom) || !g.failed.Load() {
		t.Fatal(g.err)
	}
	if g.anyPending() {
		t.Fatal("pending")
	}
	ts, err := u.targets()
	if err != nil {
		t.Fatal(err)
	}
	g = &ingester{u: u, ts: ts, budget: 1}
	ts[0].pending = true
	if !g.anyPending() {
		t.Fatal("pending")
	}
	loc := g.locals()
	if len(loc) != 2 || len(loc[0]) != 2 {
		t.Fatal(loc)
	}
	b := grammar.RecordOf(fixture.Label(1, fixture.Tok("Q1"), "en", "x"))
	d, _ := grammar.Frame(b)
	extract(ts[0], d, 1, loc[0])
	extract(ts[1], d, 1, loc[1])
	if err := g.sink(loc, 10); err != nil || g.maxDoc != 10 || g.read != 1 {
		t.Fatal(err)
	}
	ts[0].sp[0].tmp = filepath.Join(dir, currentFile, "x")
	ts[0].sp[0].buf.Add(nil, make([]byte, ts[0].sp[0].info.RowWidth))
	if err := g.flushAll(); err == nil {
		t.Fatal("run under a file")
	}
	if err := u.finishTarget(ts[0], &Report{}); err == nil {
		t.Fatal("finish under a file")
	}
	if err := u.finishTargets(ts, &Report{}); err == nil {
		t.Fatal("finish under a file")
	}
	if err := u.finishSection(&Section{Pending: true}, filepath.Join(dir, "none"), &Report{}); err == nil {
		t.Fatal("missing run")
	}
}
