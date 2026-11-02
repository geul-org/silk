//ff:func feature=cli type=builder control=sequence
//ff:what flags.indexDecl — an INDEX declaration from -part, -key and -width
package cli

import (
	"strings"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

func (f *flags) indexDecl(g *grammar.Grammar, name, tag string) (decl.Decl, error) {
	ws, err := parseWidths(f.width)
	if err != nil {
		return decl.Decl{}, err
	}
	var keys []string
	if f.key != "" {
		keys = strings.Split(f.key, ",")
	}
	return decl.NewIndex(g, name, tag, f.part, keys, ws)
}
