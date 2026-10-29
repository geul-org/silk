//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what runCond — one condition end to end: segment lookup, exact check, hits grouped by document
package query

import "github.com/geul-org/silk/pkg/index"

// runCond returns the documents meeting cp.
func runCond(ix *index.Index, cp CondPlan, threads int, st *Stats) ([]docHit, error) {
	cs, err := selectCond(ix, cp, st)
	if err != nil {
		return nil, err
	}
	ids, ok, err := checkAll(ix, cs, cp.Cond, threads)
	if err != nil {
		return nil, err
	}
	hs := groupHits(cs, ids, ok)
	for _, b := range ok {
		if b {
			st.Verified++
		}
	}
	return hs, nil
}
