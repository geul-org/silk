//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what ent — an entity from predicate/object token pairs; `?` is somevalue, `!` novalue
package sig

import "github.com/geul-org/silk/pkg/grammar"

func ent(pairs ...string) *Entity {
	e := &Entity{Objects: map[string][][]byte{}}
	for i := 0; i < len(pairs); i += 2 {
		p, _ := grammar.Ref(pairs[i])
		var o []byte
		switch pairs[i+1] {
		case "?":
			o = []byte(wordSomeValue)
		case "!":
			o = []byte(wordNoValue)
		default:
			o, _ = grammar.Ref(pairs[i+1])
		}
		e.Objects[string(p)] = append(e.Objects[string(p)], o)
	}
	return e
}
