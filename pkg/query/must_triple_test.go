//ff:func feature=query type=helper control=sequence
//ff:what mustTriple — ParseTriple that fails the test on error
package query

import "testing"

// mustTriple parses s.
func mustTriple(t *testing.T, s string) Triple {
	t.Helper()
	tr, err := ParseTriple(s)
	if err != nil {
		t.Fatal(err)
	}
	return tr
}
