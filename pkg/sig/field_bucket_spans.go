//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what Field.bucketSpans — the code of a bucket by name, or of the bucket holding a number; none outside every bucket
package sig

import (
	"fmt"
	"strconv"
)

func (f *Field) bucketSpans(v string) ([]Span, error) {
	if e := f.entry(v); e != nil {
		return []Span{{e.Code, e.Code}}, nil
	}
	y, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: bucket %s needs a name or number", ErrCond, f.Name)
	}
	var out []Span
	for _, e := range f.Codes {
		if e.Lo <= y && y < e.Hi {
			out = append(out, Span{e.Code, e.Code})
		}
	}
	return out, nil
}
