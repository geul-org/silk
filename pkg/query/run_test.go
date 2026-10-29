//ff:func feature=query type=scanner control=sequence
//ff:what TestRun — physicists in the fixture (two documents), conjunctions narrow, empty queries and scan errors fail
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/scan"
)

func TestRun(t *testing.T) {
	ix := fixtureIndex(t)
	ctx := context.Background()
	rs, st, err := Run(ctx, ix, Query{Triples: []Triple{mustTriple(t, "P106=Q169470")}}, scan.Options{})
	if err != nil || len(rs) != 2 || rs[0].Doc != 0x5100000000001d52 || rs[1].Edges[0] != 2 || st.Maps != 1 || st.Hits != 2 {
		t.Fatal(rs, st, err)
	}
	both := Query{Triples: []Triple{mustTriple(t, "P106=Q169470"), mustTriple(t, "P31=Q5")}}
	if rs, _, err := Run(ctx, ix, both, scan.Options{}); err != nil || len(rs) != 2 || len(rs[0].Edges) != 2 {
		t.Fatal(rs, err)
	}
	if _, _, err := Run(ctx, ix, Query{}, scan.Options{}); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
	cancel, stop := context.WithCancel(ctx)
	stop()
	if _, _, err := Run(cancel, ix, both, scan.Options{}); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
