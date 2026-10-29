//ff:func feature=query type=scanner control=sequence
//ff:what selectCond — a condition's candidates from its section (live versions only), sorted by document and offset
package query

import (
	"cmp"
	"slices"

	"github.com/park-jun-woo/silk/pkg/index"
)

// selectCond runs the segment lookup of one condition.
func selectCond(ix *index.Index, cp CondPlan, st *Stats) ([]candidate, error) {
	var cs []candidate
	s, err := ix.Select(cp.Section, cp.Sel, func(doc, off uint32) error {
		cs = append(cs, candidate{doc: doc, off: off})
		return nil
	})
	st.Segments += s.Segments
	st.Rows += s.Rows
	st.Dead += s.Dead
	st.Candidates += s.Hits
	slices.SortFunc(cs, func(a, b candidate) int {
		if a.doc != b.doc {
			return cmp.Compare(a.doc, b.doc)
		}
		return cmp.Compare(a.off, b.off)
	})
	return cs, err
}
