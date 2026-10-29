//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what parseQuery — the -triple, -match and -label values as one Query (in that order)
package cli

import "github.com/geul-org/silk/pkg/query"

// parseQuery reads every condition; -label L is -match LABEL:2.0=@L.
func parseQuery(f queryFlags) (query.Query, error) {
	var q query.Query
	for _, s := range f.triples {
		c, err := query.ParseTriple(s)
		if err != nil {
			return query.Query{}, err
		}
		q.Conds = append(q.Conds, c)
	}
	for _, s := range append(append([]string{}, f.matches...), labelMatches(f.labels)...) {
		c, err := query.ParseMatch(s)
		if err != nil {
			return query.Query{}, err
		}
		q.Conds = append(q.Conds, c)
	}
	return q, nil
}
