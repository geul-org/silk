//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestRun — single and joined conditions over table and other partitions give the same documents; no condition, a cancelled context and a failing condition fail
package query

import (
	"context"
	"errors"
	"slices"
	"testing"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	want := []string{"0x00010000/0x5100000000001d52/0x00000001.geul [15 5]", "0x00010000/0x5100000000001d82/0x00000001.geul [28 2]"}
	for _, min := range []int{1, 256} {
		ix := testIndex(t, min)
		p, _ := NewPlan(ix, Query{Conds: []Cond{condOf(t, "P31=Q5"), condOf(t, "P106=Q169470")}}, 2)
		rs, st, err := Run(ctx, ix, p, 2)
		if err != nil || !slices.Equal(pathsOf(rs), want) || st.Verified != 4 {
			t.Fatal(min, err, pathsOf(rs))
		}
	}
	ix := testIndex(t, 1)
	if _, _, err := Run(ctx, ix, Plan{}, 1); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
	p, _ := NewPlan(ix, Query{Conds: []Cond{condOf(t, "P31=Q5")}}, 1)
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, _, err := Run(cctx, ix, p, 1); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
	p.Conds[0].Cond.Codec = 0x0B
	if _, _, err := Run(ctx, ix, p, 1); err == nil {
		t.Fatal("failing condition")
	}
}
