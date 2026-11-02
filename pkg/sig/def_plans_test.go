//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefPlans — when codes pick their blocks, in one plan each; a code without a block keeps the common fields
package sig

import "testing"

func TestDefPlans(t *testing.T) {
	d := testDef(t)
	ps := d.plans(ent("P31", "Q5", "P31", "Q515", "P31", "Q99", "P106", "Q36180"))
	if len(ps) != 3 {
		t.Fatal(len(ps))
	}
	if n := len(ps[0].fields); n != len(d.Common) {
		t.Fatal(n)
	}
	nw, _ := Parse("silk s 8\nfield a 0 4 TRIPLE P1 enum\n  code x 4 Q1\n")
	if ps := nw.plans(ent("P1", "Q1")); len(ps) != 1 {
		t.Fatal(ps)
	}
	if d.variant(ent(), plan{fields: d.Common, codes: [][]uint64{{4}, {0}}}, d.Common[0], 4).codes[0][0] != 4 {
		t.Fatal("variant")
	}
	if c := ent("P31", "Q5").fieldCodes(d.Common); len(c) != 2 {
		t.Fatal(c)
	}
}
