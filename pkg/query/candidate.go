//ff:func feature=query type=builder control=sequence
//ff:what candidate — how one section could answer a condition; nil when it cannot
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
)

func candidate(ix *index.Index, c *Cond, name, force string) *Plan {
	d := ix.Decls[name]
	if force != "" && force != name || c.Decl != nil && c.Decl.Name != name || d.Kind == decl.KindSilk {
		return nil
	}
	if ix.Section(name, 0) == nil || !slices.Equal(d.Tags, c.Tags) {
		return nil
	}
	p := &Plan{Cond: c, Name: name, Decl: d, Prefix: [][]byte{nil}}
	slots := fillPart(ix, p)
	fillKeys(ix, p, slots)
	p.Covered = isCovered(c, slots, p)
	return p
}
