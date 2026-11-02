//ff:func feature=cli type=builder control=selection
//ff:what flags.declaration — the declaration of `create KIND NAME [TAG]`
package cli

import (
	"fmt"
	"os"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

func (f *flags) declaration(g *grammar.Grammar, kind, name string, rest []string) (decl.Decl, error) {
	if kind == decl.KindSilk {
		text, err := os.ReadFile(f.def)
		if err != nil {
			return decl.Decl{}, err
		}
		return decl.NewSilk(name, string(text))
	}
	if len(rest) != 1 {
		return decl.Decl{}, fmt.Errorf("create %s NAME TAG", kind)
	}
	switch kind {
	case decl.KindIndex:
		return f.indexDecl(g, name, rest[0])
	case decl.KindText:
		return decl.NewText(g, name, rest[0])
	case decl.KindComposite:
		text, err := os.ReadFile(f.def)
		if err != nil {
			return decl.Decl{}, err
		}
		return decl.NewComposite(g, name, rest[0], string(text))
	}
	return decl.Decl{}, fmt.Errorf("create: unknown kind %q", kind)
}
