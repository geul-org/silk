//ff:func feature=index type=helper control=sequence
//ff:what TestPickRefs — references of the named files in the files' order, unknown names skipped
package index

import "testing"

func TestPickRefs(t *testing.T) {
	rs := pickRefs([]SegRef{{File: "a", Rows: 1}, {File: "b", Rows: 2}}, []string{"b", "x", "a"})
	if len(rs) != 2 || rs[0].Rows != 2 || rs[1].Rows != 1 {
		t.Fatal(rs)
	}
}
