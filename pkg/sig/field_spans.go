//ff:func feature=sig type=decoder control=selection
//ff:what Field.spans — the codes of a value by the field's kind
package sig

import "fmt"

func (f *Field) spans(v string) ([]Span, error) {
	switch f.Kind {
	case KindFlag:
		if v != "0" && v != "1" {
			return nil, fmt.Errorf("%w: flag %s is 0 or 1", ErrCond, f.Name)
		}
		c := uint64(v[0] - '0')
		return []Span{{c, c}}, nil
	case KindBucket:
		return f.bucketSpans(v)
	}
	return f.tableSpans(v), nil
}
