//ff:func feature=index type=scanner control=sequence
//ff:what TestGone — when a pack leaves the set its documents die and a rewritten set's come back by the new pack
package index

import (
	"context"
	"errors"
	"testing"
)

func TestGone(t *testing.T) {
	ctx := context.Background()
	data, dir, a := built(t, 3)
	m := current(t, dir)
	u := &updater{ar: a, dir: dir, o: testOpts, m: m}
	if g := u.gonePacks(); len(g) != 0 {
		t.Fatal(g)
	}
	if err := u.loadLive(); err != nil {
		t.Fatal(err)
	}
	if err := u.killGone(); err != nil {
		t.Fatal(err)
	}
	u.m.Packs = append(u.m.Packs, PackRef{Set: "0x00010000", Hash: "gone"}, PackRef{Set: "nope", Hash: "x"})
	if g := u.gonePacks(); len(g) != 2 {
		t.Fatal(g)
	}
	u.m.Packs[0].Hash = "gone too"
	if err := u.killGone(); err != nil {
		t.Fatal(err)
	}
	if u.live.count() != 0 {
		t.Fatal(u.live.count())
	}
	u.m.Docs = []DocFile{{File: "doc/none", Count: 1}}
	if err := u.killGone(); err == nil {
		t.Fatal("missing table")
	}
	if _, err := u.read(ctx, DocEntry{Pack: 0}); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	_ = data
}
