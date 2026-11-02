//ff:func feature=index type=scanner control=sequence
//ff:what TestTierRuns — the lowest full tier merges; a base below 2 never does
package index

import "testing"

func TestTierRuns(t *testing.T) {
	runs := []SegRef{{File: "a", Rows: 1}, {File: "b", Rows: 2}, {File: "c", Rows: 100}, {File: "d", Rows: 90}, {File: "e", Rows: 80}}
	if r := tierRuns(runs, 2); len(r) != 3 || r[0].File != "c" {
		t.Fatal(r)
	}
	if tierRuns(runs, 1) != nil || tierRuns(runs[:1], 2) != nil {
		t.Fatal("no tier")
	}
	if w := withoutSegs(runs, runs[1:]); len(w) != 1 {
		t.Fatal(w)
	}
}
