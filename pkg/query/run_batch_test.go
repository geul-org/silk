//ff:func feature=query type=scanner control=sequence
//ff:what TestRunBatch — answers line up with queries; the first error stops
package query

import (
	"context"
	"errors"
	"testing"
)

func TestRunBatch(t *testing.T) {
	ctx := context.Background()
	_, _, ix := smallIndex(t)
	a, _ := Parse(ix, []string{"TRIPLE", "predicate=P106", "object=Q98"})
	b, _ := Parse(ix, []string{"DESCRIPTION", "subject=Q1"})
	rs, err := RunBatch(ctx, ix, []Query{a, a})
	if err != nil || len(rs) != 2 || len(rs[1].Hits) != 1 {
		t.Fatal(rs, err)
	}
	if rs, err := RunBatch(ctx, ix, []Query{a, b}); !errors.Is(err, ErrNoPlan) || len(rs) != 1 {
		t.Fatal(rs, err)
	}
}
