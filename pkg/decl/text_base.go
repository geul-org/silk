//ff:func feature=decl type=builder control=selection
//ff:what textBase — the shared part of TEXT and composite declarations; the TAG's field 2 must take text
package decl

import (
	"fmt"
	"slices"

	"github.com/geul-org/silk/pkg/grammar"
)

func textBase(g *grammar.Grammar, name, tag string) (Decl, error) {
	tags, err := Tags(g, tag)
	if err != nil {
		return Decl{}, err
	}
	fs := g.Fields[tags[0]]
	if len(fs) < 2 || len(tags) != 1 {
		return Decl{}, fmt.Errorf("%w: %s is not a text Edge", ErrDecl, tag)
	}
	d := Decl{Name: name, Tags: tags}
	o := fs[1]
	switch {
	case slices.Contains(o.Codecs, grammar.TagLangText):
		d.Part, d.PartWidth, d.Keys = o.Name+".language", LangWidth, []string{o.Name + ".bytes"}
	case slices.Contains(o.Codecs, grammar.TagText):
		d.Keys = []string{o.Name + ".bytes"}
	default:
		return Decl{}, fmt.Errorf("%w: %s takes no text", ErrDecl, tag)
	}
	d.Widths = []int{TextWidth}
	return d, nil
}
