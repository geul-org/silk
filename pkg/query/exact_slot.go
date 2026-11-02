//ff:func feature=query type=rule control=iteration dimension=1
//ff:what exactSlot — each value fits its slot whole
package query

import "github.com/geul-org/silk/pkg/grammar"

func exactSlot(q Pred, w int) bool {
	for _, v := range q.Vals {
		if !grammar.SlotExact(grammar.Val{B: v, Rest: q.Leaf.Field.Width == grammar.RestWidth}, w) {
			return false
		}
	}
	return true
}
