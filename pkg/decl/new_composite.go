//ff:func feature=decl type=builder control=sequence
//ff:what NewComposite — a text + silk declaration: a text Edge and a 32-bit definition for the meaning half of its SIDX
package decl

import (
	"fmt"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/sig"
)

func NewComposite(g *grammar.Grammar, name, tag, text string) (Decl, error) {
	def, err := sig.Parse(text)
	if err != nil {
		return Decl{}, err
	}
	if def.Bits != SIDXMeaning {
		return Decl{}, fmt.Errorf("%w: the meaning half is %d bits, not %d", ErrDecl, def.Bits, SIDXMeaning)
	}
	d, err := textBase(g, name, tag)
	d.Kind, d.Silk = KindComposite, text
	return d, err
}
