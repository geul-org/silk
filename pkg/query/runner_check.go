//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.check — open each candidate document once and keep what its records confirm
package query

func (r *runner) check(p *Plan, hits []Hit) ([]Hit, error) {
	var out []Hit
	for i := 0; i < len(hits); {
		j := i
		for j < len(hits) && hits[j].Doc == hits[i].Doc {
			j++
		}
		hs, err := r.checkDoc(p, hits[i:j])
		if err != nil {
			return nil, err
		}
		out = append(out, hs...)
		i = j
	}
	return out, nil
}
