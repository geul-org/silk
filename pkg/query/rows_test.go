//ff:func feature=query type=scanner control=sequence
//ff:what TestRows — a section's rows
package query

import "testing"

func TestRows(t *testing.T) {
	_, _, ix := smallIndex(t)
	if n := rows(ix.Section("alias", 0)); n != 1 {
		t.Fatal(n)
	}
}
