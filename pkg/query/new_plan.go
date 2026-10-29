//ff:func feature=query type=builder control=iteration dimension=1
//ff:what NewPlan — the plan of a query over an index: each condition mapped onto its section, and the stages (signature filter off, partition/sort lookup, exact check, merge join)
package query

import (
	"fmt"

	"github.com/park-jun-woo/silk/pkg/index"
)

// NewPlan lays out how Run answers q with threads checking goroutines.
func NewPlan(ix *index.Index, q Query, threads int) (Plan, error) {
	p := Plan{Query: q}
	lookup := ""
	for _, c := range q.Conds {
		cp, err := planCond(ix, c)
		if err != nil {
			return Plan{}, err
		}
		p.Conds = append(p.Conds, cp)
		lookup += fmt.Sprintf("; %s: section %s, %s", c.Text, ix.M.Sections[cp.Section].Name, describeSel(cp))
	}
	p.Steps = []Step{
		{Name: "signature filter", Detail: "skipped: the index has no entity signatures yet"},
		{Name: "segment lookup", Active: true, Detail: "binary search in sorted runs, scan of appended segments, dead versions skipped by the live bitmap" + lookup},
		{Name: "exact check", Active: true, Detail: fmt.Sprintf("each candidate decoded with gogeul core, every field compared, Edge Id read (%d threads)", threads)},
		{Name: "document join", Active: len(q.Conds) > 1, Detail: "merge join of the conditions' documents by sequence number"},
	}
	return p, nil
}
