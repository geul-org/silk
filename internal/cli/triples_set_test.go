//ff:func feature=cli type=helper control=sequence
//ff:what TestTriplesSet — each Set appends
package cli

import "testing"

func TestTriplesSet(t *testing.T) {
	var ts triples
	if ts.Set("a") != nil || ts.Set("b") != nil || len(ts) != 2 || ts[1] != "b" {
		t.Fatal(ts)
	}
}
