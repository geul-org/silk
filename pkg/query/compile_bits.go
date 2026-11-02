//ff:func feature=query type=builder control=sequence
//ff:what compileBits — a composite section's meaning conditions as patterns over its 32 bits
package query

import (
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

func compileBits(p *Plan) error {
	if p.Decl.Kind != decl.KindComposite || len(p.Cond.Silk) == 0 {
		return nil
	}
	def, err := sig.Parse(p.Decl.Silk)
	if err != nil {
		return err
	}
	p.Def = def
	p.Bits, err = def.Compile(p.Cond.Silk)
	return err
}
