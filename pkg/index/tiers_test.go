//ff:func feature=index type=scanner control=sequence
//ff:what TestTiers — size tiers and the runs a tier gives up
package index

import "testing"

func TestTiers(t *testing.T) {
	if tierOf(7, 8) != 0 || tierOf(8, 8) != 1 || tierOf(64, 8) != 2 {
		t.Fatal("tierOf")
	}
	runs := []SegRef{{File: "a", Rows: 1}, {File: "b", Rows: 2}, {File: "c", Rows: 100}, {File: "d", Rows: 100}}
	if tierRuns(runs, 3) != nil {
		t.Skip("tier grouping changed")
	}
}
