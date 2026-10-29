//ff:func feature=index type=builder control=sequence
//ff:what TestApplyResult — inputs leave runs and appended segments, a non-empty output joins the runs, dead rows drop (never below zero)
package index

import "testing"

func TestApplyResult(t *testing.T) {
	m := Manifest{Sections: []Section{{Runs: []SegRef{{File: "a"}, {File: "keep"}}, Appends: []SegRef{{File: "b"}}, Dead: 5}}}
	applyResult(&m, taskResult{Task: Task{Inputs: []SegRef{{File: "a"}, {File: "b"}}}, Out: SegRef{File: "c", Rows: 3}, Dropped: 2})
	s := m.Sections[0]
	if len(s.Runs) != 2 || s.Runs[1].File != "c" || len(s.Appends) != 0 || s.Dead != 3 {
		t.Fatal(s)
	}
	applyResult(&m, taskResult{Task: Task{Inputs: []SegRef{{File: "c"}}}, Out: SegRef{File: "d"}, Dropped: 9})
	if s := m.Sections[0]; len(s.Runs) != 1 || s.Dead != 0 {
		t.Fatal(s)
	}
}
