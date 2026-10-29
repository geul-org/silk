//ff:func feature=index type=rule control=sequence
//ff:what TestTierInputs — the smallest tier with fanout runs, never the base run; too few runs or no full tier give none
package index

import "testing"

func TestTierInputs(t *testing.T) {
	rs := []SegRef{{File: "base", Rows: 100000}, {File: "a", Rows: 20}, {File: "b", Rows: 30}, {File: "c", Rows: 200}, {File: "d", Rows: 300},
		{File: "e", Rows: 25}, {File: "f", Rows: 250}, {File: "g", Rows: 260}, {File: "h", Rows: 18}}
	in := tierInputs(rs, 4)
	if len(in) != 4 || in[0].Rows > 30 {
		t.Fatal(in)
	}
	if tierInputs(rs[:4], 4) != nil {
		t.Fatal("too few runs")
	}
	if tierInputs([]SegRef{{Rows: 1000}, {Rows: 1}, {Rows: 10}, {Rows: 100}, {Rows: 400}}, 4) != nil {
		t.Fatal("no full tier")
	}
}
