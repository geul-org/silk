//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestScoped — plain field, when code held, when code loose, when code absent
package sig

import "testing"

func TestScoped(t *testing.T) {
	d := testDef(t)
	f, w := d.field("job")
	x, _ := f.term("writer")
	cases := []struct {
		e         *Entity
		w         *When
		has, some bool
	}{
		{ent("P106", "Q36180"), nil, true, false},
		{ent("P31", "Q5", "P106", "Q36180"), w, true, false},
		{ent("P31", "Q99", "P106", "Q36180"), w, false, true},
		{ent("P31", "?", "P106", "?"), w, false, true},
		{ent("P31", "Q515", "P106", "Q36180"), w, false, false},
		{ent("P31", "Q99"), w, false, false},
	}
	for _, c := range cases {
		if h, s := scoped(c.e, c.w, f, x, "writer"); h != c.has || s != c.some {
			t.Fatal(c, h, s)
		}
	}
}
