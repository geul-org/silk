//ff:func feature=query type=builder control=iteration dimension=1
//ff:what fillKeys — extend the key prefixes field by field while predicates fix them
package query

import "github.com/geul-org/silk/pkg/index"

func fillKeys(ix *index.Index, p *Plan, slots map[int]int) {
	for k, key := range p.Decl.Keys {
		l, _ := ix.Grammar.Resolve(p.Decl.Tags[0], key)
		i := predAt(p.Cond, l.Path)
		if i < 0 {
			return
		}
		w := p.Decl.Widths[k]
		slots[i] = w
		p.Fixed++
		q := p.Cond.Preds[i]
		p.Prefix = extend(p.Prefix, q, flagged(p.Decl, l), w)
		if q.Prefix {
			return
		}
	}
}
