//ff:func feature=query type=rule control=sequence
//ff:what anyMatch — some pattern holds
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/sig"
)

func anyMatch(ps []sig.Pattern, bits []byte) bool {
	return slices.ContainsFunc(ps, func(p sig.Pattern) bool { return p.Match(bits) })
}
