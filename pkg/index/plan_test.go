//ff:func feature=index type=rule control=sequence
//ff:what TestPlan — at most one task per section, in section order, defaults applied
package index

import "testing"

func TestPlan(t *testing.T) {
	m := Manifest{Sections: []Section{
		{Name: "A", Runs: []SegRef{{File: "a", Rows: 10}}, Appends: []SegRef{{File: "b", Rows: 5}}},
		{Name: "B", Runs: []SegRef{{File: "c", Rows: 10}}},
		{Name: "C", Runs: []SegRef{{File: "d", Rows: 10}}, Dead: 9},
	}}
	ts := Plan(m, Policy{}, false)
	if len(ts) != 2 || ts[0].Name != "A" || ts[0].Kind != TaskFlush || ts[1].Section != 2 || ts[1].Kind != TaskMajor {
		t.Fatal(ts)
	}
}
