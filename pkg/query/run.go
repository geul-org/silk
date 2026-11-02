//ff:func feature=query type=builder control=iteration dimension=1
//ff:what Run — answer a query: plan each condition, narrow by rows, then check what rows cannot decide
package query

import (
	"context"
	"fmt"

	"github.com/geul-org/silk/pkg/index"
)

func Run(ctx context.Context, ix *index.Index, q Query) (Result, error) {
	r := &runner{ix: ix, ctx: ctx, seen: map[uint32]bool{}}
	if q.Scan {
		return r.scan(q)
	}
	var res Result
	var hits []Hit
	for i := range q.Conds {
		p, err := makePlan(ix, &q.Conds[i], q.Index)
		if err != nil {
			return res, err
		}
		res.Plans = append(res.Plans, p.describe())
		if q.Mode == ModeKeys && !p.keysOK() {
			return res, fmt.Errorf("%w: %s", ErrNotCovered, p.Name)
		}
		if hits, err = r.stage(i, p, hits, q); err != nil {
			return res, err
		}
	}
	res.Hits, res.Decoded = trimBare(hits), len(r.seen)
	return res, nil
}
