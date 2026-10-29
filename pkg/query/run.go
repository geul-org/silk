//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what Run — follow a plan: each condition's documents, joined by document, as Results in (Issuer, document, version) order
package query

import (
	"context"
	"fmt"

	"github.com/park-jun-woo/silk/pkg/index"
)

// Run answers p over ix.
func Run(ctx context.Context, ix *index.Index, p Plan, threads int) ([]Result, Stats, error) {
	if len(p.Conds) == 0 {
		return nil, Stats{}, fmt.Errorf("%w: no condition", ErrSyntax)
	}
	var acc []docHit
	var st Stats
	for i, cp := range p.Conds {
		if err := ctx.Err(); err != nil {
			return nil, st, err
		}
		hs, err := runCond(ix, cp, threads, &st)
		if err != nil {
			return nil, st, err
		}
		if i == 0 {
			acc = hs
		} else {
			acc = joinDocs(acc, hs)
		}
	}
	return toResults(ix, acc), st, nil
}
