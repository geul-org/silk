//ff:func feature=grammar type=decoder control=sequence
//ff:what Grammar.field — field i of rec: 0 is an Edge's Edge Id, declared fields follow; inside a Value or Reference fields count from 0
package grammar

func (g *Grammar) field(rec []byte, i int) (Val, bool) {
	t := Target(rec)
	fs, ok := g.Fields[t]
	if !ok {
		return Val{}, false
	}
	q := headLen(rec)
	if t>>11 == kindEdge {
		n := RecLen(rec[q:])
		if n <= 0 || n > len(rec)-q {
			return Val{}, false
		}
		if i == 0 {
			return Val{B: rec[q : q+n]}, true
		}
		q += n
		i--
	}
	if i < 0 || i >= len(fs) {
		return Val{}, false
	}
	return g.cut(rec, t, fs, i, q)
}
