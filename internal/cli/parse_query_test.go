//ff:func feature=cli type=helper control=sequence
//ff:what TestParseQuery — triples, matches and labels in that order; a bad triple or match fails
package cli

import "testing"

func TestParseQuery(t *testing.T) {
	q, err := parseQuery(queryFlags{triples: triples{"P31=Q5"}, matches: triples{"SOURCE:2=#1"}, labels: triples{"ko"}})
	if err != nil || len(q.Conds) != 3 || q.Conds[1].Name != "SOURCE" || q.Conds[2].Name != "LABEL" {
		t.Fatal(err, q)
	}
	if _, err := parseQuery(queryFlags{triples: triples{"x"}}); err == nil {
		t.Fatal("bad triple")
	}
	if _, err := parseQuery(queryFlags{matches: triples{"x"}}); err == nil {
		t.Fatal("bad match")
	}
}
