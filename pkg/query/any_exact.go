//ff:func feature=query type=rule control=sequence
//ff:what anyExact — some pattern holds through its exact codes
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/sig"
)

func anyExact(ps []sig.Pattern, row []byte) bool {
	return slices.ContainsFunc(ps, func(p sig.Pattern) bool { return p.Exact(row) })
}
