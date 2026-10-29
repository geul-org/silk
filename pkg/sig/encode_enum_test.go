//ff:func feature=sig type=encoder control=sequence
//ff:what TestEncodeEnum — listed value gets its code, unlisted gets Other, none writes nothing, two values fail
package sig

import (
	"errors"
	"testing"
)

func TestEncodeEnum(t *testing.T) {
	f := Field{Bit: 2, Width: 2, Codes: map[uint64]uint64{7: 1}, Other: 3}
	s := make([]uint64, 1)
	encodeEnum(s, f, []uint64{7})
	if s[0] != 1<<2 {
		t.Fatal(s)
	}
	s[0] = 0
	encodeEnum(s, f, []uint64{8})
	if s[0] != 3<<2 {
		t.Fatal(s)
	}
	s[0] = 0
	if err := encodeEnum(s, f, nil); err != nil || s[0] != 0 {
		t.Fatal(s, err)
	}
	if err := encodeEnum(s, f, []uint64{7, 8}); !errors.Is(err, ErrValue) {
		t.Fatal(err)
	}
}
