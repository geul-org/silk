//ff:func feature=sig type=builder control=selection
//ff:what Field.reservedTerm — the term of one reserved code
package sig

func (f *Field) reservedTerm(c uint64) Term {
	s := []Span{{c, c}}
	t := Term{Off: f.Off, Width: f.Width, Spans: s, Exact: s, Truth: s}
	switch c {
	case CodeOther:
		t.Exact = nil
	case CodeNoValue, CodeSomeValue:
		t.Spans = []Span{{c, c}, {CodeOther, CodeOther}}
	}
	return t
}
