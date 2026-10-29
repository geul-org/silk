//ff:func feature=query type=builder control=iteration dimension=1
//ff:what toResults — joined documents as Results sorted by Issuer, document ID, version
package query

import (
	"sort"

	"github.com/park-jun-woo/silk/pkg/index"
)

// toResults names each document through the index's document table.
func toResults(ix *index.Index, hs []docHit) []Result {
	out := make([]Result, 0, len(hs))
	for _, h := range hs {
		d := ix.Docs[h.doc]
		out = append(out, Result{Issuer: ix.M.Issuers[d.Issuer], Doc: d.ID, Version: d.Version, Edges: h.edges})
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Issuer != b.Issuer {
			return a.Issuer < b.Issuer
		}
		if a.Doc != b.Doc {
			return a.Doc < b.Doc
		}
		return a.Version < b.Version
	})
	return out
}
