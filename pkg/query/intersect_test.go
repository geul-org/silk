//ff:func feature=query type=builder control=sequence
//ff:what TestIntersect — the first condition starts the set; later ones keep shared documents and append Edge Ids
package query

import (
	"slices"
	"testing"
)

func TestIntersect(t *testing.T) {
	a, b := docKey{doc: 1}, docKey{doc: 2}
	first := intersect(nil, map[docKey][]uint64{a: {1}, b: {2}}, true)
	got := intersect(first, map[docKey][]uint64{b: {5}}, false)
	if len(got) != 1 || !slices.Equal(got[b], []uint64{2, 5}) {
		t.Fatal(got)
	}
}
