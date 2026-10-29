//ff:func feature=sig type=helper control=sequence
//ff:what TestMask — bits across a word boundary
package sig

import (
	"slices"
	"testing"
)

func TestMask(t *testing.T) {
	if got := mask(62, 4, 2); !slices.Equal(got, []uint64{3 << 62, 3}) {
		t.Fatal(got)
	}
}
