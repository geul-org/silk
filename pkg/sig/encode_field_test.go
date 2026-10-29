//ff:func feature=sig type=encoder control=sequence
//ff:what TestEncodeField — enum and one-hot write; Bloom, Prefix and Bucket are not implemented yet
package sig

import (
	"errors"
	"testing"
)

func TestEncodeField(t *testing.T) {
	s := make([]uint64, 1)
	if err := encodeField(s, Field{Kind: Enum, Width: 2, Codes: map[uint64]uint64{7: 2}}, []uint64{7}); err != nil || s[0] != 2 {
		t.Fatal(s, err)
	}
	if err := encodeField(s, Field{Kind: OneHot, Bit: 4, Width: 2, Bits: map[uint64]int{7: 1}}, []uint64{7}); err != nil || s[0] != 2|1<<5 {
		t.Fatal(s, err)
	}
	if err := encodeField(s, Field{Kind: Bloom}, nil); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
}
