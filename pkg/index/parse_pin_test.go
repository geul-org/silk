//ff:func feature=index type=decoder control=iteration dimension=1
//ff:what TestParsePin — generation and process ID; wrong field count or non-numbers are false
package index

import "testing"

func TestParsePin(t *testing.T) {
	if g, p, ok := parsePin("4-99-123"); !ok || g != 4 || p != 99 {
		t.Fatal(g, p, ok)
	}
	for _, s := range []string{"4-99", "x-1-2", "1-y-2"} {
		if _, _, ok := parsePin(s); ok {
			t.Fatal(s)
		}
	}
}
