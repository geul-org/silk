//ff:func feature=grammar type=decoder control=selection
//ff:what Grammar.lookup — a part by name or number in one codec's layout; 0 of an Edge is its Edge Id
package grammar

import "strconv"

func (g *Grammar) lookup(t uint16, part string) (int, Field, bool) {
	fs, ok := g.Fields[t]
	if !ok {
		return 0, Field{}, false
	}
	edge := 0
	if t>>11 == kindEdge {
		edge = 1
	}
	n, err := strconv.Atoi(part)
	switch {
	case err == nil && edge == 1 && n == 0:
		return 0, Field{Name: "id"}, true
	case err == nil && n-edge >= 0 && n-edge < len(fs):
		return n, fs[n-edge], true
	case err == nil:
		return 0, Field{}, false
	}
	return lookupName(fs, part, edge)
}
