//ff:func feature=query type=scanner control=sequence
//ff:what TestSilkFilter — hits leave when no row of their document meets the patterns, stay on an exact row, and go to the document on a candidate row
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
)

func TestSilkFilter(t *testing.T) {
	ctx := context.Background()
	_, _, ix := smallIndex(t)
	all := []Hit{{Doc: 0}, {Doc: 1}}
	r := &runner{ix: ix, ctx: ctx, seen: map[uint32]bool{}}
	ps := func(w ...string) *Plan {
		q, err := Parse(ix, append([]string{"entity"}, w...))
		if err != nil {
			t.Fatal(err)
		}
		p, err := makePlan(ix, &q.Conds[0], "")
		if err != nil {
			t.Fatal(err)
		}
		return p
	}
	keep := ix.Manifest.Packs[0].Hash
	ix.Manifest.Packs[0].Hash = "gone"
	if _, err := r.silkFilter(ps("job=Q777"), all, Query{}); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	b := []Hit{{Doc: 1}, {Doc: 1, HasOff: true}}
	if _, err := r.checkDoc(&Plan{Decl: decl.Decl{}}, b); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	if _, err := r.check(&Plan{}, b); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	ix.Manifest.Packs[0].Hash = keep
	if h, err := r.silkFilter(ps("job=physicist"), all, Query{}); err != nil || len(h) != 1 || h[0].Doc != 0 {
		t.Fatal(h, err)
	}
	if h, err := r.silkFilter(ps("job=politician"), all, Query{}); err != nil || len(h) != 0 {
		t.Fatal(h, err)
	}
	if h, err := r.silkFilter(ps("type=taxon"), all, Query{}); err != nil || len(h) != 0 {
		t.Fatal(h, err)
	}
}
