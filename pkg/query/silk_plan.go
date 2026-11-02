//ff:func feature=query type=builder control=sequence
//ff:what silkPlan — a condition on a silk: its patterns in value order
package query

import "github.com/geul-org/silk/pkg/sig"

func silkPlan(c *Cond) (*Plan, error) {
	def, err := sig.Parse(c.Decl.Silk)
	if err != nil {
		return nil, err
	}
	ps, err := def.Compile(c.Silk)
	if err != nil {
		return nil, err
	}
	return &Plan{Cond: c, Name: c.Name, Decl: *c.Decl, Pats: ps, Def: def, Covered: true}, nil
}
