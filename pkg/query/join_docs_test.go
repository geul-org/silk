//ff:func feature=query type=builder control=sequence
//ff:what TestJoinDocs — documents in both lists, Edge Ids of the first then the second
package query

import (
	"slices"
	"testing"
)

func TestJoinDocs(t *testing.T) {
	a := []docHit{{1, []uint64{1}}, {3, []uint64{3}}, {5, []uint64{5}}}
	b := []docHit{{2, []uint64{20}}, {3, []uint64{30}}, {5, []uint64{50}}, {7, nil}}
	j := joinDocs(a, b)
	if len(j) != 2 || j[0].doc != 3 || !slices.Equal(j[1].edges, []uint64{5, 50}) || len(a[1].edges) != 1 {
		t.Fatal(j)
	}
}
