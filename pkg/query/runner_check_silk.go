//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.checkSilk — exact rows answer; candidates (somevalue, other) go to their documents and stay only if the statements agree
package query

func (r *runner) checkSilk(p *Plan, hits []Hit) ([]Hit, error) {
	best := map[uint32]bool{}
	for _, h := range hits {
		if !h.Loose {
			best[h.Doc] = true
		}
	}
	var out []Hit
	for _, h := range hits {
		if !h.Loose {
			out = append(out, Hit{Doc: h.Doc, Key: h.Key})
			continue
		}
		if best[h.Doc] {
			continue
		}
		yes, unknown, err := r.holds(p, h.Doc)
		if err != nil {
			return nil, err
		}
		if yes || unknown {
			out = append(out, Hit{Doc: h.Doc, Key: h.Key, Unknown: !yes})
			best[h.Doc] = true
		}
	}
	return uniqueDocs(out), nil
}
