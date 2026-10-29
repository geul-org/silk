//ff:func feature=sig type=encoder control=sequence
//ff:what TestSchemaPattern — enum equality and one-hot membership combine; unknown fields and values fail
package sig

import (
	"errors"
	"slices"
	"testing"
)

func TestSchemaPattern(t *testing.T) {
	p, err := sampleSchema().Pattern(map[string]uint64{"P31": 5, "P106": 169470})
	if err != nil || !slices.Equal(p.Mask, []uint64{0xF | 1<<60, 0}) || !slices.Equal(p.Want, []uint64{1 | 1<<60, 0}) {
		t.Fatalf("%x %x %v", p.Mask, p.Want, err)
	}
	if _, err := sampleSchema().Pattern(map[string]uint64{"P999": 1}); !errors.Is(err, ErrValue) {
		t.Fatal(err)
	}
	if _, err := sampleSchema().Pattern(map[string]uint64{"P31": 999}); !errors.Is(err, ErrValue) {
		t.Fatal(err)
	}
}
