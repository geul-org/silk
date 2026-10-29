//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what TestSchemaValidate — the sample schema is valid; bad width, duplicate names, placement, overlap and tables fail
package sig

import (
	"errors"
	"testing"
)

func TestSchemaValidate(t *testing.T) {
	if err := sampleSchema().Validate(); err != nil {
		t.Fatal(err)
	}
	bad := []Schema{
		{Width: 8},
		{Width: 48},
		{Width: 16, Fields: []Field{{Name: "a", Kind: OneHot, Width: 2}, {Name: "a", Kind: OneHot, Bit: 4, Width: 2}}},
		{Width: 16, Fields: []Field{{Name: "a", Kind: OneHot, Bit: 15, Width: 2}}},
		{Width: 16, Fields: []Field{{Name: "a", Kind: OneHot, Width: 4}, {Name: "b", Kind: OneHot, Bit: 3, Width: 2}}},
		{Width: 16, Fields: []Field{{Name: "a", Kind: Enum, Width: 2, Codes: map[uint64]uint64{1: 4}}}},
	}
	for i, s := range bad {
		if err := s.Validate(); !errors.Is(err, ErrSchema) {
			t.Fatal(i, err)
		}
	}
}
