//ff:func feature=sig type=rule control=sequence
//ff:what Entry.clashes — same bits, two value-bearing codes on one full code, or two overlapping ranges
package sig

func (e *Entry) clashes(o *Entry, width int) bool {
	a, _ := e.span(width)
	b, _ := o.span(width)
	same := a == b && (e.Len == o.Len || len(e.Values) > 0 && len(o.Values) > 0)
	return same || e.Lo < o.Hi && o.Lo < e.Hi
}
