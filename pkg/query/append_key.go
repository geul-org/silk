//ff:func feature=query type=encoder control=sequence
//ff:what appendKey — one prefix extended by one value
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/grammar"
)

func appendKey(base []byte, q Pred, rest bool, v []byte, w int) []byte {
	b := slices.Clone(base)
	if q.Prefix {
		return append(b, v[:min(len(v), w)]...)
	}
	return grammar.AppendSlot(b, grammar.Val{B: v, Rest: rest}, true, w)
}
