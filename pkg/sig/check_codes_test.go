//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what TestCheckCodes — codes 1…2^w-1 pass (64-bit fields too); zero, overflow and a too-large Other fail
package sig

import (
	"errors"
	"testing"
)

func TestCheckCodes(t *testing.T) {
	if err := checkCodes(Field{Width: 2, Codes: map[uint64]uint64{9: 3}, Other: 3}); err != nil {
		t.Fatal(err)
	}
	if err := checkCodes(Field{Width: 64, Codes: map[uint64]uint64{9: ^uint64(0)}}); err != nil {
		t.Fatal(err)
	}
	for _, f := range []Field{{Width: 2, Other: 4}, {Width: 2, Codes: map[uint64]uint64{9: 0}}, {Width: 2, Codes: map[uint64]uint64{9: 4}}} {
		if err := checkCodes(f); !errors.Is(err, ErrSchema) {
			t.Fatal(f, err)
		}
	}
}
