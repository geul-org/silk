//ff:func feature=cli type=helper control=sequence
//ff:what TestTriplesString — values joined by spaces
package cli

import "testing"

func TestTriplesString(t *testing.T) {
	ts := triples{"P1=Q2", "P3=*"}
	if ts.String() != "P1=Q2 P3=*" {
		t.Fatal(ts.String())
	}
}
