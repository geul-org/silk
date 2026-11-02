//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.scanDoc — one document against every condition
package query

func (r *runner) scanDoc(seq uint32, q Query) ([]Hit, error) {
	d, err := r.doc(seq)
	if err != nil {
		return nil, err
	}
	var hits []Hit
	for i := range q.Conds {
		hs, err := r.scanCond(seq, d, &q.Conds[i])
		if err != nil || len(hs) == 0 {
			return nil, err
		}
		hits = append(hits, hs...)
	}
	return trimBare(sortHits(hits)), nil
}
