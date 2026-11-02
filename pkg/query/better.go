//ff:func feature=query type=rule control=sequence
//ff:what better — more fixed key fields, then fewer rows, then a narrower row
package query

import "github.com/geul-org/silk/pkg/index"

func better(ix *index.Index, p, best *Plan) bool {
	if best == nil {
		return true
	}
	if p.Fixed != best.Fixed {
		return p.Fixed > best.Fixed
	}
	a, b := ix.Section(p.Name, 0), ix.Section(best.Name, 0)
	if x, y := rows(a), rows(b); x != y {
		return x < y
	}
	return p.Name < best.Name
}
