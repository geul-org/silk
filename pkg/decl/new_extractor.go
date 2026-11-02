//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what NewExtractor — compile a declaration against its grammar
package decl

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/sig"
)

func NewExtractor(d Decl, g *grammar.Grammar) (*Extractor, error) {
	x := &Extractor{d: d, g: g, tags: map[uint16]bool{}, sub: grammar.Path{1}}
	for _, t := range d.Tags {
		x.tags[t] = true
	}
	if err := x.compilePaths(); err != nil {
		return nil, err
	}
	if d.Silk == "" {
		return x, nil
	}
	def, err := sig.Parse(d.Silk)
	x.def = def
	return x, err
}
