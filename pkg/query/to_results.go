//ff:func feature=query type=builder control=iteration dimension=1
//ff:what toResults — collected documents as Results sorted by Issuer, document ID, version
package query

import "sort"

// toResults orders the run's output deterministically.
func toResults(acc map[docKey][]uint64) []Result {
	out := make([]Result, 0, len(acc))
	for k, e := range acc {
		out = append(out, Result{Issuer: k.issuer, Doc: k.doc, Version: k.version, Edges: e})
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
