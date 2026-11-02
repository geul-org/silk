//ff:func feature=query type=builder control=sequence
//ff:what runner.stage — one condition's hits, joined by document with those so far; a later silk condition reads the document-order bundle
package query

func (r *runner) stage(i int, p *Plan, prev []Hit, q Query) ([]Hit, error) {
	if i > 0 && p.Pats != nil {
		return r.silkFilter(p, prev, q)
	}
	hits, err := r.selectRows(p, q)
	if err != nil {
		return nil, err
	}
	if i == 0 {
		return hits, nil
	}
	return joinDocs(prev, hits), nil
}
