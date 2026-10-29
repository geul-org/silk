//ff:func feature=index type=decoder control=sequence
//ff:what TestParsePairs — pairs back to TagCounts
package index

import "testing"

func TestParsePairs(t *testing.T) {
	cs := parsePairs(appendPairs(nil, []TagCount{{1, 2}, {3, 4}}))
	if len(cs) != 2 || cs[1] != (TagCount{3, 4}) {
		t.Fatal(cs)
	}
}
