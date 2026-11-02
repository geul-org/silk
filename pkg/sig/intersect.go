//ff:func feature=sig type=builder control=sequence
//ff:what intersect — the spans two terms on one field share
package sig

func intersect(a, b Term) Term {
	return Term{Off: a.Off, Width: a.Width, Spans: cross(a.Spans, b.Spans), Exact: cross(a.Exact, b.Exact)}
}
