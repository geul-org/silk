//ff:func feature=query type=builder control=iteration dimension=1
//ff:what makePlan — the best section for a condition: a silk condition its own section, else the longest key prefix; none is ErrNoPlan with the index to create
package query

import (
	"fmt"
	"maps"
	"slices"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
)

func makePlan(ix *index.Index, c *Cond, force string) (*Plan, error) {
	if c.Decl != nil && c.Decl.Kind == decl.KindSilk {
		return silkPlan(c)
	}
	var best *Plan
	for _, name := range slices.Sorted(maps.Keys(ix.Decls)) {
		p := candidate(ix, c, name, force)
		if p != nil && better(ix, p, best) {
			best = p
		}
	}
	if best == nil {
		return nil, fmt.Errorf("%w: %s; try: %s", ErrNoPlan, c.Name, suggest(c))
	}
	return best, compileBits(best)
}
