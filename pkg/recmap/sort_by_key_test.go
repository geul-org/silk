//ff:func feature=recmap type=builder control=sequence
//ff:what TestSortByKey — offsets grouped by ascending key, ascending within a key; ranges tile the result
package recmap

import (
	"slices"
	"testing"
)

func TestSortByKey(t *testing.T) {
	ranges, sorted := sortByKey([]uint16{5, 2, 5, 1}, []uint32{10, 20, 30, 40})
	want := []KeyRange{{Key: 1, Start: 0, Count: 1}, {Key: 2, Start: 1, Count: 1}, {Key: 5, Start: 2, Count: 2}}
	if !slices.Equal(ranges, want) || !slices.Equal(sorted, []uint32{40, 20, 10, 30}) {
		t.Fatal(ranges, sorted)
	}
}
