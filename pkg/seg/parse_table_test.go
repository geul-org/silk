//ff:func feature=seg type=decoder control=iteration dimension=1
//ff:what TestParseTable — contiguous ascending entries parse; a gap, an empty or oversized range, descending keys or uncovered rows fail
package seg

import (
	"errors"
	"testing"
)

func TestParseTable(t *testing.T) {
	info := Info{PartWidth: 1, Rows: 5}
	mk := func(ps ...Part) []byte { return appendTable(nil, ps) }
	if ps, err := parseTable(mk(Part{[]byte("a"), 0, 2}, Part{[]byte("b"), 2, 3}), info); err != nil || len(ps) != 2 {
		t.Fatal(err)
	}
	bad := [][]byte{
		mk(Part{[]byte("a"), 1, 4}),
		mk(Part{[]byte("a"), 0, 0}),
		mk(Part{[]byte("a"), 0, 6}),
		mk(Part{[]byte("b"), 0, 2}, Part{[]byte("a"), 2, 3}),
		mk(Part{[]byte("a"), 0, 2}),
	}
	for i, b := range bad {
		if _, err := parseTable(b, info); !errors.Is(err, ErrFormat) {
			t.Fatalf("case %d: %v", i, err)
		}
	}
}
