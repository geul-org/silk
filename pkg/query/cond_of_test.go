//ff:func feature=query type=helper control=sequence
//ff:what condOf — test helper: a condition from text (KIND:… is a match, else a triple), failing the test on a parse error
package query

import (
	"strings"
	"testing"
)

// condOf parses s.
func condOf(t testing.TB, s string) Cond {
	t.Helper()
	parse := ParseTriple
	if strings.Contains(s, ":") {
		parse = ParseMatch
	}
	c, err := parse(s)
	if err != nil {
		t.Fatal(err)
	}
	return c
}
