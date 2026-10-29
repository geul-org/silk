//ff:func feature=index type=helper control=sequence
//ff:what TestCountOf — a kind's count, 0 when absent
package index

import "testing"

func TestCountOf(t *testing.T) {
	cs := []TagCount{{1, 4}, {2, 5}}
	if countOf(cs, 2) != 5 || countOf(cs, 9) != 0 {
		t.Fatal("countOf")
	}
}
