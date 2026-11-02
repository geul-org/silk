//ff:func feature=query type=rule control=iteration dimension=1
//ff:what rowsMeet — whether a document's rows meet the patterns exactly, and whether any meets them as a candidate
package query

import (
	"github.com/geul-org/silk/pkg/seg"
	"github.com/geul-org/silk/pkg/sig"
)

func rowsMeet(ps []sig.Pattern, reg seg.Region) (bool, bool) {
	exact, loose := false, false
	for i := 0; i < reg.Len(); i++ {
		bits := reg.Row(i)[4:]
		exact = exact || anyExact(ps, bits)
		loose = loose || anyMatch(ps, bits)
	}
	return exact, loose
}
