//ff:func feature=query type=rule control=iteration dimension=1
//ff:what isCovered — every predicate was used by the key and each slot holds its whole value
package query

import "github.com/geul-org/silk/pkg/decl"

func isCovered(c *Cond, slots map[int]int, p *Plan) bool {
	if len(slots) != len(c.Preds) || p.Decl.Kind != decl.KindIndex {
		return false
	}
	for i, q := range c.Preds {
		if q.Prefix || !exactSlot(q, slots[i]) {
			return false
		}
	}
	return true
}
