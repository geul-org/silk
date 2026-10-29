//ff:func feature=cli type=helper control=sequence
//ff:what TestParseQuery — every condition parsed in order; one bad condition fails
package cli

import "testing"

func TestParseQuery(t *testing.T) {
	q, err := parseQuery([]string{"P106=Q169470", "P31=Q5"})
	if err != nil || len(q.Triples) != 2 || q.Triples[1].Text != "P31=Q5" {
		t.Fatal(q, err)
	}
	if _, err := parseQuery([]string{"P1=Q2", "bad"}); err == nil {
		t.Fatal("bad")
	}
}
