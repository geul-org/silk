//ff:func feature=query type=scanner control=sequence
//ff:what TestEdgeID — a hit's record, its Edge Id, and hits without offsets
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/index"
)

func TestEdgeID(t *testing.T) {
	ctx := context.Background()
	a, dir, ix := smallIndex(t)
	q, _ := Parse(ix, []string{"TRIPLE", "predicate=P106", "object=Q98"})
	r, err := Run(ctx, ix, q)
	if err != nil || len(r.Hits) != 1 {
		t.Fatal(r, err)
	}
	n, ok, err := EdgeID(ctx, ix, r.Hits[0])
	if err != nil || !ok || n != 2 {
		t.Fatal(n, ok, err)
	}
	if rec, err := Record(ctx, ix, r.Hits[0]); err != nil || grammar.Target(rec) != 0x100A {
		t.Fatal(rec, err)
	}
	if rec, err := Record(ctx, ix, Hit{}); rec != nil || err != nil {
		t.Fatal(rec, err)
	}
	if _, ok, err := EdgeID(ctx, ix, Hit{}); ok || err != nil {
		t.Fatal(ok, err)
	}
	if _, err := Record(ctx, ix, Hit{HasOff: true, Off: 1 << 30}); !errors.Is(err, index.ErrFormat) {
		t.Fatal(err)
	}
	b, _ := ix.Bytes(ctx, 0)
	d, _ := grammar.Frame(b)
	last := d.Recs[len(d.Recs)-1]
	if _, ok, err := EdgeID(ctx, ix, Hit{HasOff: true, Off: uint32(last.Off), Doc: 0}); ok || err != nil {
		t.Fatal("END has no Edge Id", ok, err)
	}
	cold, err := index.Open(ctx, a, dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	if _, _, err := EdgeID(ctx, cold, Hit{HasOff: true, Doc: 1}); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	if _, err := Record(ctx, cold, Hit{HasOff: true, Doc: 1}); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
}
