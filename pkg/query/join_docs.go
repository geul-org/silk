//ff:func feature=query type=helper control=iteration dimension=1
//ff:what joinDocs — hits of a whose document is also in b, plus those of b
package query

func joinDocs(a, b []Hit) []Hit {
	in := func(hs []Hit) map[uint32]bool {
		m := map[uint32]bool{}
		for _, h := range hs {
			m[h.Doc] = true
		}
		return m
	}
	ma, mb := in(a), in(b)
	var out []Hit
	for _, h := range a {
		if mb[h.Doc] {
			out = append(out, h)
		}
	}
	for _, h := range b {
		if ma[h.Doc] {
			out = append(out, h)
		}
	}
	return sortHits(out)
}
