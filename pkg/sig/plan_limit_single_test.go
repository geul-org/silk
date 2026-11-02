//ff:func feature=sig type=scanner control=sequence
//ff:what TestPlanLimitSingle — a one-code field is left alone
package sig

import "testing"

func TestPlanLimitSingle(t *testing.T) {
	a := &Field{Width: 4}
	p := plan{fields: []*Field{a, a}, codes: [][]uint64{{4}, {5, 6}}}
	if l := p.limit(1); len(l.codes[1]) != 2 || len(l.codes[0]) != 1 {
		t.Fatal(l.codes)
	}
}
