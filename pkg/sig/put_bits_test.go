//ff:func feature=sig type=encoder control=sequence
//ff:what TestPutBits — the low bits of a value land at the offset, across a word boundary
package sig

import (
	"slices"
	"testing"
)

func TestPutBits(t *testing.T) {
	s := make([]uint64, 2)
	putBits(s, 62, 4, 0b1011)
	if !slices.Equal(s, []uint64{3 << 62, 2}) {
		t.Fatal(s)
	}
}
