//ff:func feature=sig type=helper control=sequence
//ff:what newParser — a parser over an empty definition
package sig

func newParser() *parser {
	return &parser{d: &Def{Preds: map[string]bool{}, Bits: 16}}
}
