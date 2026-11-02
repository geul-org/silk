//ff:func feature=sig type=builder control=sequence
//ff:what Field.term — the term of `field=value`: its codes plus somevalue and other as candidates; a reserved name asks for that code, keeping `other` as a candidate for somevalue and novalue since a limited row may hide them, and `other` itself is never exact since limited rows write it too
package sig

func (f *Field) term(v string) (Term, error) {
	if c, ok := reservedCode(v); ok && f.Kind != KindFlag {
		return f.reservedTerm(c), nil
	}
	s, err := f.spans(v)
	if err != nil {
		return Term{}, err
	}
	if f.Kind == KindFlag {
		return Term{Off: f.Off, Width: f.Width, Spans: s, Exact: s, Truth: s}, nil
	}
	all := append([]Span{{CodeSomeValue, CodeOther}}, s...)
	return Term{Off: f.Off, Width: f.Width, Spans: all, Exact: s, Truth: s}, nil
}
