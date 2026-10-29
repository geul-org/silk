//ff:func feature=query type=builder control=iteration dimension=1
//ff:what joinDocs — merge join of two document-sorted hit lists: documents in both, Edge Ids of the first then the second
package query

// joinDocs intersects acc with hits in one pass.
func joinDocs(acc, hits []docHit) []docHit {
	var out []docHit
	i, j := 0, 0
	for i < len(acc) && j < len(hits) {
		switch {
		case acc[i].doc < hits[j].doc:
			i++
		case acc[i].doc > hits[j].doc:
			j++
		default:
			out = append(out, docHit{doc: acc[i].doc, edges: append(append([]uint64{}, acc[i].edges...), hits[j].edges...)})
			i, j = i+1, j+1
		}
	}
	return out
}
