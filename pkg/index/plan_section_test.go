//ff:func feature=index type=rule control=iteration dimension=1
//ff:what TestPlanSection — empty: none; full, too many segments or too many dead rows: major; many appended rows: flush; a full size tier: tier; else none
package index

import "testing"

func TestPlanSection(t *testing.T) {
	p := DefaultPolicy()
	base := []SegRef{{File: "base", Rows: 1000}}
	runs := func(n int, rows uint64) []SegRef {
		out := append([]SegRef{}, base...)
		for i := range n {
			out = append(out, SegRef{File: string(rune('a' + i)), Rows: rows})
		}
		return out
	}
	cases := []struct {
		s    Section
		full bool
		kind string
	}{
		{Section{}, false, ""},
		{Section{Runs: base}, true, ""},
		{Section{Runs: base, Appends: []SegRef{{File: "x", Rows: 1}}}, true, TaskMajor},
		{Section{Runs: runs(8, 1)}, false, TaskMajor},
		{Section{Runs: base, Dead: 400}, false, TaskMajor},
		{Section{Runs: base, Appends: []SegRef{{File: "x", Rows: 200}}}, false, TaskFlush},
		{Section{Runs: base, Appends: []SegRef{{File: "x", Rows: 10}}}, false, ""},
		{Section{Runs: runs(4, 20)}, false, TaskTier},
		{Section{Runs: runs(3, 20)}, false, ""},
	}
	for i, c := range cases {
		task, ok := planSection(0, c.s, p, c.full)
		if ok != (c.kind != "") || task.Kind != c.kind {
			t.Fatal(i, task, ok)
		}
	}
}
