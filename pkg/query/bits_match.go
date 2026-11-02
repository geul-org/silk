//ff:func feature=query type=rule control=sequence
//ff:what bitsMatch — a composite row's meaning bits meet the condition's patterns
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

func bitsMatch(p *Plan, row []byte) bool {
	if p.Bits == nil {
		return true
	}
	bits := row[decl.TextWidth : decl.TextWidth+4]
	return slices.ContainsFunc(p.Bits, func(s sig.Pattern) bool { return s.Match(bits) })
}
