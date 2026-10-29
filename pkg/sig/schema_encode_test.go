//ff:func feature=sig type=encoder control=sequence
//ff:what TestSchemaEncode — enum codes and one-hot bits land in place; missing fields stay 0; a field error fails
package sig

import (
	"errors"
	"slices"
	"testing"
)

func TestSchemaEncode(t *testing.T) {
	s, err := sampleSchema().Encode(map[string][]uint64{"P31": {5}, "P106": {169470, 36180}})
	if err != nil || !slices.Equal(s, []uint64{1 | 1<<60, 1 << 3}) {
		t.Fatalf("%x %v", s, err)
	}
	if _, err := sampleSchema().Encode(map[string][]uint64{"P31": {5, 515}}); !errors.Is(err, ErrValue) {
		t.Fatal(err)
	}
}
