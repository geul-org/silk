//ff:func feature=sig type=builder control=sequence
//ff:what Def.exactTerm — the when field holds this block's code; a condition on that field itself is narrower or equal, so it is dropped by merging
package sig

func (d *Def) exactTerm(w *When) Term {
	lo, hi := w.Entry.span(w.On.Width)
	return Term{Off: w.On.Off, Width: w.On.Width, Spans: []Span{{lo, hi}}, Exact: []Span{{lo, hi}}}
}
