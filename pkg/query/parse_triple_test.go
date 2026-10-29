//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what TestParseTriple — Predicate (field 1) and Object (field 3) conditions, wildcards dropped; no '=', bad tokens or two wildcards fail
package query

import (
	"errors"
	"slices"
	"testing"
)

func TestParseTriple(t *testing.T) {
	c, err := ParseTriple("P106 = Q169470")
	if err != nil || c.Name != "TRIPLE" || c.Codec != 0x0A || len(c.Fields) != 2 || !slices.Equal(c.Fields[1].Path, []int{3}) {
		t.Fatal(c, err)
	}
	if c, _ := ParseTriple("*=Q5"); len(c.Fields) != 1 || c.Fields[0].Path[0] != 3 {
		t.Fatal(c)
	}
	for _, s := range []string{"P106", "X=Q5", "P1=X", "*=*"} {
		if _, err := ParseTriple(s); !errors.Is(err, ErrSyntax) {
			t.Fatal(s, err)
		}
	}
}
