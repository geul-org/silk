//ff:func feature=query type=rule control=sequence
//ff:what runner.scanSilk — the silk conditions read from a document's statements
package query

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/sig"
)

func (r *runner) scanSilk(seq uint32, d *grammar.Doc, c *Cond) (bool, error) {
	def, err := sig.Parse(c.Decl.Silk)
	if err != nil {
		return false, err
	}
	return r.entitiesHold(def, d, c.Silk)
}
