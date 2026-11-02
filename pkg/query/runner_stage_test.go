//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerStage — the first condition selects, a later one joins, a later silk filters
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
)

func TestRunnerStage(t *testing.T) {
	ar, dir, ix := smallIndex(t)
	r := testRunner(ix)
	cold, _ := index.Open(context.Background(), ar, dir, 0)
	defer cold.Close()
	cold.Manifest.Packs[0].Hash = "gone"
	lq, _ := Parse(cold, []string{"LABEL", `object.bytes="Short"`})
	lp, _ := makePlan(cold, &lq.Conds[0], "")
	if _, err := testRunner(cold).stage(0, lp, nil, lq); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	q, _ := Parse(ix, []string{"TRIPLE", "predicate=P31", "object=Q99", "TRIPLE", "predicate=P106", "object=Q98", "entity", "type=taxon"})
	a, _ := makePlan(ix, &q.Conds[0], "")
	hs, err := r.stage(0, a, nil, q)
	if err != nil || len(hs) != 1 {
		t.Fatal(hs, err)
	}
	b, _ := makePlan(ix, &q.Conds[1], "")
	if hs, err = r.stage(1, b, hs, q); err != nil || len(hs) != 2 {
		t.Fatal(hs, err)
	}
	c, _ := makePlan(ix, &q.Conds[2], "")
	if hs, err = r.stage(2, c, hs, q); err != nil || len(hs) != 0 {
		t.Fatal(hs, err)
	}
	d, _ := makePlan(ix, &Cond{Tags: []uint16{0x1004}, Name: "LABEL"}, "")
	d.Decl.Kind = decl.KindText
	d.Decl.Silk = ""
	dead := ix.Live
	ix.Live = index.Bitmap{}
	if _, err := r.stage(1, d, []Hit{{Doc: 0}}, q); err != nil {
		t.Fatal(err)
	}
	ix.Live = dead
}
