//ff:func feature=index type=helper control=sequence
//ff:what TestAddCount — a new kind is appended, a known one counted up
package index

import "testing"

func TestAddCount(t *testing.T) {
	cs := addCount(addCount(addCount(nil, 1), 2), 1)
	if len(cs) != 2 || cs[0] != (TagCount{1, 2}) || cs[1] != (TagCount{2, 1}) {
		t.Fatal(cs)
	}
}
