//ff:func feature=sig type=scanner control=sequence
//ff:what TestPlanLimit — keep 1 leaves the first multi field, keep 0 none; a flag keeps set; emit crosses what stays
package sig

import (
	"slices"
	"testing"
)

func TestPlanLimit(t *testing.T) {
	a := &Field{Name: "a", Off: 0, Width: 4}
	b := &Field{Name: "b", Off: 4, Width: 1, Kind: KindFlag}
	p := plan{fields: []*Field{a, b}, codes: [][]uint64{{4, 5}, {0, 1}}}
	if l := p.limit(1); !slices.Equal(l.codes[0], []uint64{4, 5}) || !slices.Equal(l.codes[1], []uint64{1}) {
		t.Fatal(l.codes)
	}
	if l := p.limit(0); !slices.Equal(l.codes[0], []uint64{CodeOther}) {
		t.Fatal(l.codes)
	}
	if n := len(p.emit(nil, []byte{0}, 0)); n != 4 {
		t.Fatal(n)
	}
	if wildcard(b) != 1 || wildcard(a) != CodeOther {
		t.Fatal("wildcard")
	}
	if q := p.bounded(2, 1); &q.codes[0][0] != &p.codes[0][0] {
		t.Fatal("under the bound")
	}
	if q := p.bounded(9, 1); len(q.codes[0]) != 2 || len(q.codes[1]) != 1 {
		t.Fatal(q.codes)
	}
	if q := p.bounded(9, 2); len(q.codes[0]) != 1 {
		t.Fatal(q.codes)
	}
}
