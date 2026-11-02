//ff:func feature=index type=scanner control=sequence
//ff:what TestDocStops — a document that does not frame stops the workers and the feed
package index

import (
	"context"
	"errors"
	"testing"
)

func TestDocStops(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	g := &ingester{u: u}
	j := job{seq: 1, entry: DocEntry{Pack: 99}}
	func() {
		defer func() { recover() }()
		g.doc(ctx, j, nil)
	}()
	u.m.Packs[0].Hash = "gone"
	if err := g.doc(ctx, job{}, nil); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	jobs := make(chan job, 2)
	jobs <- job{}
	jobs <- job{}
	close(jobs)
	g.worker(ctx, jobs)
	if !g.failed.Load() {
		t.Fatal("worker error")
	}
	u.live.Put(0, true)
	u.live.Put(1, false)
	u.live.Put(2, true)
	docs := &DocTable{}
	u.m.Docs = nil
	if err := g.feed(ctx, docs, make(chan job, 1)); err != nil {
		t.Fatal(err)
	}
}
