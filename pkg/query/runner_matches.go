//ff:func feature=query type=rule control=iteration dimension=1
//ff:what runner.matches — every predicate holds on a record
package query

import "github.com/geul-org/silk/pkg/grammar"

func (r *runner) matches(d *grammar.Doc, rec []byte, c *Cond) bool {
	for _, q := range c.Preds {
		v, ok := d.Get(r.ix.Grammar, rec, q.Leaf.Path)
		if !ok || !anyEqual(q, v.B) {
			return false
		}
	}
	return true
}
