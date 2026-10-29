//ff:func feature=index type=helper control=sequence
//ff:what TestAddDead — a removed version's records of each section's kind count as dead, except in rebuilt sections
package index

import "testing"

func TestAddDead(t *testing.T) {
	secs := []Section{{Tag: 1}, {Tag: 2}, {Tag: 3}}
	addDead(secs, []TagCount{{1, 4}, {2, 5}}, []bool{false, true, false})
	if secs[0].Dead != 4 || secs[1].Dead != 0 || secs[2].Dead != 0 {
		t.Fatal(secs)
	}
}
