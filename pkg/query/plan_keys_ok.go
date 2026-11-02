//ff:func feature=query type=rule control=sequence
//ff:what Plan.keysOK — -keys answers from rows: a covered index, or a text or silk section whose rows are the answer
package query

import "github.com/geul-org/silk/pkg/decl"

func (p *Plan) keysOK() bool {
	return p.Covered || p.Decl.Kind != decl.KindIndex
}
