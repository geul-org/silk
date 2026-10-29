//ff:func feature=index type=helper control=iteration dimension=1
//ff:what byIssuer — documents grouped by Issuer (offset order kept) and the Issuers sorted
package index

import (
	"sort"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// byIssuer groups docs; each group keeps docs' order.
func byIssuer(docs []archive.Doc) ([]string, map[string][]archive.Doc) {
	by := map[string][]archive.Doc{}
	var names []string
	for _, d := range docs {
		if _, ok := by[d.Issuer]; !ok {
			names = append(names, d.Issuer)
		}
		by[d.Issuer] = append(by[d.Issuer], d)
	}
	sort.Strings(names)
	return names, by
}
