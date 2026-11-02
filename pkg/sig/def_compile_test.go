//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestDefCompile — common, when and contradictory conditions; bad fields and values fail
package sig

import (
	"errors"
	"testing"
)

func TestDefCompile(t *testing.T) {
	d := testDef(t)
	ps, err := d.Compile([]Cond{{"job", "writer"}})
	if err != nil || len(ps) != 2 {
		t.Fatal(ps, err)
	}
	row := make([]byte, 8)
	put(row, 0, 6, 4)
	put(row, 7, 8, 5)
	if !ps[0].Match(row) || !ps[0].Exact(row) {
		t.Fatal("person writer")
	}
	put(row, 7, 8, CodeOther)
	if !ps[0].Match(row) || ps[0].Exact(row) {
		t.Fatal("other is a candidate")
	}
	put(row, 0, 6, CodeOther)
	put(row, 7, 8, 0)
	if ps[0].Match(row) || !ps[1].Match(row) {
		t.Fatal("type other stays a candidate")
	}
	ps, err = d.Compile([]Cond{{"type", "person"}, {"job", "writer"}, {"born", "1950"}})
	if err != nil || len(ps) != 2 || len(ps[0].Terms) != 3 {
		t.Fatal(ps, err)
	}
	ps, err = d.Compile([]Cond{{"job", "writer"}, {"pop", "big"}})
	if err != nil || len(ps) != 1 {
		t.Fatal(ps, err)
	}
	ps, err = d.Compile([]Cond{{"alive", "1"}})
	if err != nil || len(ps) != 1 || len(ps[0].Terms) != 1 {
		t.Fatal(ps, err)
	}
	for _, c := range []Cond{{"nope", "1"}, {"alive", "2"}, {"born", "x"}} {
		if _, err := d.Compile([]Cond{c}); !errors.Is(err, ErrCond) {
			t.Fatal(c, err)
		}
	}
	if x := cross([]Span{{1, 5}, {8, 9}}, []Span{{4, 8}}); len(x) != 2 || x[1].Lo != 8 {
		t.Fatal(x)
	}
	if len(overlaps(Span{1, 2}, []Span{{3, 4}})) != 0 {
		t.Fatal("overlaps")
	}
	it := intersect(Term{Off: 1, Spans: []Span{{1, 2}}, Exact: []Span{{1, 1}}}, Term{Spans: []Span{{2, 3}}})
	if len(it.Spans) != 1 || it.Exact != nil {
		t.Fatal(it)
	}
	if w := d.whenOf(d.Common[0], 9); w != nil {
		t.Fatal(w)
	}
}
