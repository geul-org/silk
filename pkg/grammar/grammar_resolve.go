//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what Grammar.Resolve — a dotted path of field numbers or names under a target word, checked against the declared layouts
package grammar

import (
	"fmt"
	"strings"
)

func (g *Grammar) Resolve(tag uint16, s string) (Leaf, error) {
	ts := []uint16{tag}
	var l Leaf
	for _, part := range strings.Split(s, ".") {
		i, f, next, err := g.step(ts, part)
		if err != nil {
			return Leaf{}, fmt.Errorf("%s in %q: %w", part, s, err)
		}
		l.Path, l.Field, ts = append(l.Path, i), f, next
	}
	return l, nil
}
