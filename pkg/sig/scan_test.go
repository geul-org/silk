//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestScan — physicists among 10,000 encoded entities, same rows with 1, default and many threads; an empty table has none
package sig

import (
	"slices"
	"testing"
)

func TestScan(t *testing.T) {
	s := sampleSchema()
	tb := NewTable(s)
	var want []uint32
	for i := range 10000 {
		occ := []uint64{82955}
		if i%97 == 0 {
			occ = append(occ, 169470)
			want = append(want, uint32(i))
		}
		sg, _ := s.Encode(map[string][]uint64{"P31": {5}, "P106": occ})
		tb.Append(sg)
	}
	p, _ := s.Pattern(map[string]uint64{"P31": 5, "P106": 169470})
	if got := Scan(tb, p, 1); !slices.Equal(got, want) {
		t.Fatal(len(got), len(want))
	}
	if got := Scan(tb, p, 0); !slices.Equal(got, want) {
		t.Fatal(len(got))
	}
	if got := Scan(tb, p, 64); !slices.Equal(got, want) {
		t.Fatal(len(got))
	}
	if got := Scan(NewTable(s), p, 4); len(got) != 0 {
		t.Fatal(got)
	}
}
