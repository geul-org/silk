//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefPattern — terms sort by bit; two on one field intersect
package sig

import "testing"

func TestDefPattern(t *testing.T) {
	d := testDef(t)
	p := d.pattern([]Term{{Off: 7, Spans: []Span{{1, 1}}}, {Off: 0, Spans: []Span{{1, 5}}}, {Off: 0, Spans: []Span{{4, 9}}}})
	if len(p.Terms) != 2 || p.Terms[0].Spans[0] != (Span{4, 5}) || p.Terms[1].Off != 7 {
		t.Fatal(p)
	}
}
