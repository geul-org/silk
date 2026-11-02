//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.selectRows — the rows of a plan in every segment, checked unless they decide alone; -keys never opens a document
package query

func (r *runner) selectRows(p *Plan, q Query) ([]Hit, error) {
	var hits []Hit
	for _, f := range r.ix.Section(p.Name, 0).Files {
		hits = append(hits, r.fileRows(p, f)...)
	}
	hits = sortHits(hits)
	if q.Mode == ModeKeys || p.Covered && p.Pats == nil {
		return markLoose(p, hits), nil
	}
	if p.Pats != nil {
		return r.checkSilk(p, hits)
	}
	return r.check(p, hits)
}
