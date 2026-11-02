//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what Def.pattern — terms in bit order, several terms on one field intersected
package sig

import "slices"

func (d *Def) pattern(ts []Term) Pattern {
	slices.SortStableFunc(ts, func(a, b Term) int { return a.Off - b.Off })
	var out []Term
	for _, t := range ts {
		if n := len(out); n > 0 && out[n-1].Off == t.Off {
			out[n-1] = intersect(out[n-1], t)
			continue
		}
		out = append(out, t)
	}
	return Pattern{Bits: d.Bits, Terms: out}
}
