//ff:func feature=query type=rule control=iteration dimension=1
//ff:what runner.entitiesHold — any entity of a document meets or may meet the conditions
package query

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/sig"
)

func (r *runner) entitiesHold(def *sig.Def, d *grammar.Doc, cs []sig.Cond) (bool, error) {
	for _, e := range def.Entities(r.ix.Grammar, d) {
		y, u, err := def.Holds(e, cs)
		if err != nil {
			return false, err
		}
		if y || u {
			return true, nil
		}
	}
	return false, nil
}
