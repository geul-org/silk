//ff:func feature=query type=builder control=iteration dimension=1
//ff:what fillPart — the partition keys a predicate fixes; none fixed means every partition. The result maps each predicate used to its slot width
package query

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/index"
)

func fillPart(ix *index.Index, p *Plan) map[int]int {
	slots := map[int]int{}
	if p.Decl.Part == "" {
		return slots
	}
	l, _ := ix.Grammar.Resolve(p.Decl.Tags[0], p.Decl.Part)
	i := predAt(p.Cond, l.Path)
	if i < 0 {
		return slots
	}
	for _, v := range p.Cond.Preds[i].Vals {
		s := grammar.AppendSlot(nil, grammar.Val{B: v}, true, p.Decl.PartWidth)
		p.Parts = append(p.Parts, s)
	}
	slots[i] = p.Decl.PartWidth
	p.Fixed++
	return slots
}
