//ff:func feature=query type=rule control=sequence
//ff:what runner.meaningHolds — a composite plan's meaning conditions hold on the document's statements; rows hold only candidates for somevalue and other
package query

import "github.com/geul-org/silk/pkg/grammar"

func (r *runner) meaningHolds(p *Plan, d *grammar.Doc) (bool, error) {
	if p.Bits == nil {
		return true, nil
	}
	return r.entitiesHold(p.Def, d, p.Cond.Silk)
}
