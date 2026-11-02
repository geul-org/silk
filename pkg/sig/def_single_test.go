//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefSingle — one row always: a multi-valued field is other, a single when code keeps its block
package sig

import "testing"

func TestDefSingle(t *testing.T) {
	d := testDef(t)
	job, _ := d.field("job")
	r := d.Single(ent("P31", "Q5", "P106", "Q36180", "P106", "Q82955"))
	if get(r, 0, 6) != 4 || get(r, job.Off, 8) != CodeOther {
		t.Fatal(r)
	}
	r = d.Single(ent("P31", "Q5", "P31", "Q515"))
	if get(r, 0, 6) != CodeOther {
		t.Fatal(r)
	}
	nw, _ := Parse("silk s 8\nfield a 0 4 TRIPLE P1 enum\n  code x 4 Q1\n")
	if r := nw.Single(ent("P1", "Q1")); get(r, 0, 4) != 4 {
		t.Fatal(r)
	}
}
