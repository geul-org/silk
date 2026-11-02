//ff:func feature=sig type=scanner control=sequence
//ff:what TestEntityStated — the value is there, on another predicate, or no token at all
package sig

import "testing"

func TestEntityStated(t *testing.T) {
	d := testDef(t)
	job, _ := d.field("job")
	e := ent("P106", "Q7", "P21", "Q8")
	if !e.stated(job, "Q7") || e.stated(job, "Q8") || e.stated(job, "writer") {
		t.Fatal("stated")
	}
}
