//ff:func feature=sig type=encoder control=sequence
//ff:what TestEncodeOneHot — each listed value sets its bit; unlisted values are ignored
package sig

import "testing"

func TestEncodeOneHot(t *testing.T) {
	s := make([]uint64, 1)
	encodeOneHot(s, Field{Bit: 8, Width: 4, Bits: map[uint64]int{1: 0, 2: 3}}, []uint64{2, 9, 1})
	if s[0] != 1<<8|1<<11 {
		t.Fatal(s)
	}
}
