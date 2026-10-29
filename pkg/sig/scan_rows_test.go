//ff:func feature=sig type=scanner control=sequence
//ff:what TestScanRows — rows in [lo, hi) that match, by row number
package sig

import (
	"slices"
	"testing"
)

func TestScanRows(t *testing.T) {
	tb := &Table{Words: 1, Data: []uint64{1, 3, 1, 2, 1}}
	p := Pattern{Mask: []uint64{1}, Want: []uint64{1}}
	if got := scanRows(tb, p, 1, 5); !slices.Equal(got, []uint32{1, 2, 4}) {
		t.Fatal(got)
	}
}
