//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what Run — scan each condition over every map, keep the documents meeting all of them, return results in document order
package query

import (
	"context"
	"fmt"

	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/scan"
)

// Run follows the plan for q over ix.
func Run(ctx context.Context, ix *index.Index, q Query, o scan.Options) ([]Result, Stats, error) {
	if len(q.Triples) == 0 {
		return nil, Stats{}, fmt.Errorf("%w: no condition", ErrSyntax)
	}
	var acc map[docKey][]uint64
	st := Stats{Maps: len(ix.Parts)}
	for i, t := range q.Triples {
		hits, err := scanTriple(ctx, ix, t, o, &st)
		if err != nil {
			return nil, st, err
		}
		acc = intersect(acc, hits, i == 0)
	}
	return toResults(acc), st, nil
}
