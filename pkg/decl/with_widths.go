//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what withWidths — the key widths: given, or each field's default
package decl

import (
	"fmt"

	"github.com/geul-org/silk/pkg/grammar"
)

func withWidths(d Decl, g *grammar.Grammar, tag uint16, keys []string, widths []int) (Decl, error) {
	if len(widths) != 0 && len(widths) != len(keys) {
		return Decl{}, fmt.Errorf("%w: %d widths for %d keys", ErrDecl, len(widths), len(keys))
	}
	for i, k := range keys {
		l, err := g.Resolve(tag, k)
		if err != nil {
			return Decl{}, err
		}
		w := l.Width()
		if len(widths) != 0 {
			w = widths[i]
		}
		if w <= 0 {
			return Decl{}, fmt.Errorf("%w: width %d", ErrDecl, w)
		}
		d.Widths = append(d.Widths, w)
	}
	return d, nil
}
