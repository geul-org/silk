//ff:func feature=decl type=builder control=sequence
//ff:what NewIndex — an INDEX declaration: TAG name or `0x` word, partition and key paths checked against the layouts, default widths from the field kinds
package decl

import (
	"fmt"

	"github.com/geul-org/silk/pkg/grammar"
)

func NewIndex(g *grammar.Grammar, name, tag, part string, keys []string, widths []int) (Decl, error) {
	tags, err := Tags(g, tag)
	if err != nil {
		return Decl{}, err
	}
	d := Decl{Name: name, Kind: KindIndex, Tags: tags, Part: part, Keys: keys}
	if part != "" {
		l, err := g.Resolve(tags[0], part)
		if err != nil {
			return Decl{}, err
		}
		if l.Field.Width == grammar.RestWidth {
			return Decl{}, fmt.Errorf("%w: partition field %s has no fixed length", ErrDecl, part)
		}
		d.PartWidth = l.Width()
	}
	return withWidths(d, g, tags[0], keys, widths)
}
