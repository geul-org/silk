//ff:func feature=decl type=builder control=sequence
//ff:what NewText — a TEXT declaration on BIND, LABEL or ALIAS: partitioned by language when the Edge's object carries one
package decl

import "github.com/geul-org/silk/pkg/grammar"

func NewText(g *grammar.Grammar, name, tag string) (Decl, error) {
	d, err := textBase(g, name, tag)
	d.Kind = KindText
	return d, err
}
