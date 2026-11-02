//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what plan.limit — keep several codes in the first multi-valued field only (or none); the others become `other`, which any value condition keeps as a candidate
package sig

import "slices"

func (p plan) limit(keep int) plan {
	out := plan{fields: p.fields, codes: slices.Clone(p.codes)}
	seen := 0
	for i, c := range p.codes {
		if len(c) < 2 {
			continue
		}
		seen++
		if seen > keep {
			out.codes[i] = []uint64{wildcard(p.fields[i])}
		}
	}
	return out
}
