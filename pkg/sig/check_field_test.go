//ff:func feature=sig type=rule control=sequence
//ff:what TestCheckField — enum and one-hot tables are checked, declared-only kinds pass, an unknown kind fails
package sig

import (
	"errors"
	"testing"
)

func TestCheckField(t *testing.T) {
	if err := checkField(Field{Kind: Enum, Width: 2, Codes: map[uint64]uint64{1: 3}}); err != nil {
		t.Fatal(err)
	}
	if err := checkField(Field{Kind: OneHot, Width: 2, Bits: map[uint64]int{1: 2}}); !errors.Is(err, ErrSchema) {
		t.Fatal(err)
	}
	if err := checkField(Field{Kind: Bucket}); err != nil {
		t.Fatal(err)
	}
	if err := checkField(Field{Kind: 99}); !errors.Is(err, ErrSchema) {
		t.Fatal(err)
	}
}
