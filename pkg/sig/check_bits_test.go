//ff:func feature=sig type=rule control=sequence
//ff:what TestCheckBits — bits inside the field pass; negative or past the width fail
package sig

import (
	"errors"
	"testing"
)

func TestCheckBits(t *testing.T) {
	if err := checkBits(Field{Width: 3, Bits: map[uint64]int{1: 0, 2: 2}}); err != nil {
		t.Fatal(err)
	}
	if err := checkBits(Field{Width: 3, Bits: map[uint64]int{1: -1}}); !errors.Is(err, ErrSchema) {
		t.Fatal(err)
	}
	if err := checkBits(Field{Width: 3, Bits: map[uint64]int{1: 3}}); !errors.Is(err, ErrSchema) {
		t.Fatal(err)
	}
}
