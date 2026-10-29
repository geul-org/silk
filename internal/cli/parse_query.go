//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what parseQuery — the -triple values as a Query
package cli

import "github.com/park-jun-woo/silk/pkg/query"

// parseQuery reads every condition.
func parseQuery(ts []string) (query.Query, error) {
	var q query.Query
	for _, s := range ts {
		t, err := query.ParseTriple(s)
		if err != nil {
			return query.Query{}, err
		}
		q.Triples = append(q.Triples, t)
	}
	return q, nil
}
