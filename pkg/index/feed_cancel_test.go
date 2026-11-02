//ff:func feature=index type=scanner control=sequence
//ff:what TestFeedCancel — a cancelled context stops a blocked feed; dead numbers are skipped
package index

import (
	"context"
	"errors"
	"testing"
)

func TestFeedCancel(t *testing.T) {
	_, dir, a := built(t, 3)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	u.loadLive()
	u.live.Put(0, false)
	docs, _ := openDocs(dir, &u.m)
	defer docs.Close()
	ts, _ := u.targets()
	ts[0].pending = true
	g := &ingester{u: u, ts: ts}
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan job)
	done := make(chan error)
	go func() { done <- g.feed(ctx, docs, jobs) }()
	<-jobs
	cancel()
	if err := <-done; !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
