//ff:func feature=query type=rule control=sequence
//ff:what anyEqual — a value equals, or starts with for a prefix predicate, one of the predicate's byte strings
package query

import (
	"bytes"
	"slices"
)

func anyEqual(q Pred, v []byte) bool {
	return slices.ContainsFunc(q.Vals, func(w []byte) bool {
		if q.Prefix {
			return bytes.HasPrefix(v, w)
		}
		return bytes.Equal(v, w)
	})
}
