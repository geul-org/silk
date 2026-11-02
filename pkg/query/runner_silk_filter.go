//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.silkFilter — keep hits whose document has a row meeting the patterns, read from the document-order bundle
package query

import "fmt"

func (r *runner) silkFilter(p *Plan, prev []Hit, q Query) ([]Hit, error) {
	sec := r.ix.Section(p.Name, 1)
	if sec == nil {
		return nil, fmt.Errorf("%w: %s", ErrNoPlan, p.Name)
	}
	var out []Hit
	for _, h := range prev {
		ok, err := r.docHas(p, sec, h.Doc, q.Mode == ModeKeys)
		if err != nil {
			return nil, err
		}
		if ok {
			out = append(out, h)
		}
	}
	return out, nil
}
