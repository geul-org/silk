//ff:func feature=query type=builder control=iteration dimension=1
//ff:what groupHits — matched candidates (sorted by document) as one docHit per document with its Edge Ids in record order
package query

// groupHits keeps candidates whose ok flag is set.
func groupHits(cs []candidate, ids []uint64, ok []bool) []docHit {
	var out []docHit
	for i, c := range cs {
		if !ok[i] {
			continue
		}
		if n := len(out); n > 0 && out[n-1].doc == c.doc {
			out[n-1].edges = append(out[n-1].edges, ids[i])
			continue
		}
		out = append(out, docHit{doc: c.doc, edges: []uint64{ids[i]}})
	}
	return out
}
