//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.scan — the oracle: read every live document and apply every condition, no section touched
package query

func (r *runner) scan(q Query) (Result, error) {
	var res Result
	for seq := uint32(0); seq < r.ix.Live.N; seq++ {
		if !r.ix.Live.Get(seq) {
			continue
		}
		hs, err := r.scanDoc(seq, q)
		if err != nil {
			return res, err
		}
		res.Hits = append(res.Hits, hs...)
	}
	res.Plans = []string{"full scan"}
	res.Decoded = len(r.seen)
	return res, nil
}
