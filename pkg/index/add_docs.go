//ff:func feature=index type=builder control=iteration dimension=1
//ff:what addDocs — index new document versions with sequence numbers from first (every tag gets rows); their table entries
package index

// addDocs ingests docs in order.
func addDocs(g *ingest, docs []liveDoc, first uint32) ([]DocEntry, error) {
	all := make([]bool, len(g.sp.Tags))
	for i := range all {
		all[i] = true
	}
	out := make([]DocEntry, 0, len(docs))
	for i, d := range docs {
		counts, err := g.doc(d, first+uint32(i), all)
		if err != nil {
			return nil, err
		}
		out = append(out, g.entry(d, counts))
	}
	return out, nil
}
