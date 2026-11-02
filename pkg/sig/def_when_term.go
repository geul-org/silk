//ff:func feature=sig type=builder control=sequence
//ff:what Def.whenTerm — the when field is somevalue or other
package sig

func (d *Def) whenTerm(on *Field, _ bool) Term {
	return Term{Off: on.Off, Width: on.Width, Spans: []Span{{CodeSomeValue, CodeOther}}}
}
