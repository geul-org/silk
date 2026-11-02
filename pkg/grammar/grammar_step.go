//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what Grammar.step — one path part among the possible target words: its field number, field and the next target words
package grammar

import "fmt"

func (g *Grammar) step(ts []uint16, part string) (int, Field, []uint16, error) {
	for _, t := range ts {
		if i, f, ok := g.lookup(t, part); ok {
			return i, f, f.Codecs, nil
		}
	}
	return 0, Field{}, nil, fmt.Errorf("%w: no such field", ErrSyntax)
}
