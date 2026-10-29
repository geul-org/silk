//ff:func feature=sig type=encoder control=sequence
//ff:what TestPatternField — enum and one-hot conditions set mask and want; unknown values and unsupported kinds fail
package sig

import (
	"errors"
	"testing"
)

func TestPatternField(t *testing.T) {
	p := Pattern{Mask: make([]uint64, 1), Want: make([]uint64, 1)}
	if err := patternField(p, Field{Kind: Enum, Bit: 1, Width: 2, Codes: map[uint64]uint64{7: 2}}, 7); err != nil || p.Mask[0] != 6 || p.Want[0] != 4 {
		t.Fatal(p, err)
	}
	if err := patternField(p, Field{Kind: OneHot, Bit: 8, Width: 4, Bits: map[uint64]int{7: 2}}, 7); err != nil || p.Mask[0] != 6|1<<10 || p.Want[0] != 4|1<<10 {
		t.Fatal(p, err)
	}
	if err := patternField(p, Field{Kind: Enum}, 7); !errors.Is(err, ErrValue) {
		t.Fatal(err)
	}
	if err := patternField(p, Field{Kind: OneHot}, 7); !errors.Is(err, ErrValue) {
		t.Fatal(err)
	}
	if err := patternField(p, Field{Kind: Prefix}, 7); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
}
