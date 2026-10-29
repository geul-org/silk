//ff:func feature=sig type=helper control=sequence
//ff:what TestOrFree — free bits are claimed; an overlap leaves used unchanged
package sig

import (
	"slices"
	"testing"
)

func TestOrFree(t *testing.T) {
	used := []uint64{1, 0}
	if !orFree(used, []uint64{2, 4}) || !slices.Equal(used, []uint64{3, 4}) {
		t.Fatal(used)
	}
	if orFree(used, []uint64{0, 4}) || !slices.Equal(used, []uint64{3, 4}) {
		t.Fatal(used)
	}
}
