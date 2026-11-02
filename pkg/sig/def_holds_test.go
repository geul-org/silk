//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestDefHolds — yes, no and unknown on statements; a when field needs its code; somevalue and other leave it unknown; bad conditions fail
package sig

import (
	"errors"
	"testing"
)

func TestDefHolds(t *testing.T) {
	d := testDef(t)
	cases := []struct {
		e       *Entity
		c       []Cond
		yes, un bool
	}{
		{ent("P31", "Q5", "P106", "Q36180"), []Cond{{"job", "writer"}}, true, false},
		{ent("P31", "Q5", "P106", "Q82955"), []Cond{{"job", "writer"}}, false, false},
		{ent("P31", "Q5", "P106", "Q777"), []Cond{{"job", "Q777"}}, true, false},
		{ent("P31", "Q5", "P106", "?"), []Cond{{"job", "writer"}}, false, true},
		{ent("P31", "Q515", "P106", "Q36180"), []Cond{{"job", "writer"}}, false, false},
		{ent("P31", "Q99", "P106", "Q36180"), []Cond{{"job", "writer"}}, false, true},
		{ent("P31", "?", "P106", "Q36180"), []Cond{{"job", "writer"}}, false, true},
		{ent("P31", "Q99"), []Cond{{"job", "writer"}}, false, false},
		{ent("P31", "Q5"), []Cond{{"type", "person"}, {"alive", "1"}}, false, false},
		{ent("P31", "?"), []Cond{{"type", "person"}}, false, true},
	}
	for _, c := range cases {
		y, u, err := d.Holds(c.e, c.c)
		if err != nil || y != c.yes || u != c.un {
			t.Fatal(c.c, y, u, err)
		}
	}
	if _, _, err := d.Holds(ent(), []Cond{{"nope", "1"}}); !errors.Is(err, ErrCond) {
		t.Fatal(err)
	}
	if _, _, err := d.Holds(ent(), []Cond{{"alive", "9"}}); !errors.Is(err, ErrCond) {
		t.Fatal(err)
	}
	h, s := fieldHolds([]uint64{CodeSomeValue}, Term{Spans: []Span{{2, 3}}, Truth: []Span{{5, 5}}})
	if h || !s {
		t.Fatal(h, s)
	}
}
