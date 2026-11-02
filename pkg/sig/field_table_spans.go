//ff:func feature=sig type=decoder control=sequence
//ff:what Field.tableSpans — the span of a code by name or value; none for a value the definition does not list (only somevalue and other can answer)
package sig

import "github.com/geul-org/silk/pkg/grammar"

func (f *Field) tableSpans(v string) []Span {
	e := f.entry(v)
	if e == nil {
		if b, err := grammar.Ref(v); err == nil {
			e = f.byValue[string(b)]
		}
	}
	if e == nil {
		return nil
	}
	lo, hi := e.span(f.Width)
	return []Span{{lo, hi}}
}
