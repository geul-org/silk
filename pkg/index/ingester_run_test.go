//ff:func feature=index type=scanner control=sequence
//ff:what TestIngesterRun — workers read the wanted documents; a bad table entry stops them with its error
package index

import (
	"context"
	"errors"
	"testing"
)

func TestIngesterRun(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	u.loadLive()
	docs, _ := openDocs(dir, &u.m)
	defer docs.Close()
	ts, _ := u.targets()
	ts[0].pending = true
	g := &ingester{u: u, ts: ts, budget: 1 << 20}
	if err := g.run(ctx, docs); err != nil || g.read != 2 {
		t.Fatal(g.read, err)
	}
	u.m.Packs[0].Hash = "gone"
	g = &ingester{u: u, ts: ts, budget: 1 << 20}
	if err := g.run(ctx, docs); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
}
