//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefRows — spreading: one row per value up to MaxExpand; past it only one field spreads and the rest is other; bits lie where declared
package sig

import "testing"

func TestDefRows(t *testing.T) {
	d := testDef(t)
	job, _ := d.field("job")
	area, _ := d.field("area")
	rows := d.Rows(ent("P31", "Q5", "P106", "Q36180", "P106", "Q82955"))
	if len(rows) != 2 || get(rows[0], 0, 6) != 4 || get(rows[0], job.Off, 8) != 4 || get(rows[1], job.Off, 8) != 5 {
		t.Fatal(rows)
	}
	rows = d.Rows(ent("P31", "Q5", "P106", "Q36180", "P106", "Q82955", "P27", "Q884", "P27", "Q17"))
	if len(rows) != 2 || get(rows[0], area.Off, 6) != CodeOther {
		t.Fatal(rows)
	}
	rows = d.Rows(ent("P31", "Q5", "P31", "Q515", "P106", "Q36180", "P106", "Q82955", "P106", "Q9", "P1082", "Q1"))
	if len(rows) != 2 || get(rows[0], job.Off, 8) != CodeOther {
		t.Fatal(rows)
	}
	rows = d.Rows(ent("P31", "Q99"))
	if len(rows) != 1 || get(rows[0], 0, 6) != CodeOther {
		t.Fatal(rows)
	}
	nw, _ := Parse("silk s 8\nfield a 0 4 TRIPLE P1 enum\n  code x 4 Q1\n")
	if r := nw.Rows(ent("P1", "Q1")); len(r) != 1 || nw.whenField() != nil {
		t.Fatal(r)
	}
	if p := (plan{codes: [][]uint64{{1, 2}, {3}}}).product(); p != 2 {
		t.Fatal(p)
	}
}
