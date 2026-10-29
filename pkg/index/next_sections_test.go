//ff:func feature=index type=builder control=sequence
//ff:what TestNextSections — rebuilt tags get fresh sections with their runs, kept tags keep runs and dead rows and gain appended segments; dropped tags vanish
package index

import (
	"testing"
)

func TestNextSections(t *testing.T) {
	sp := smallSpec()
	sp.Tags = sp.Tags[:2]
	old := []Section{{Name: "QUALIFIER", Tag: 0x100B, Runs: []SegRef{{File: "q"}}, Dead: 3}, {Name: "GONE"}}
	out := [][]SegRef{{{File: "t1"}, {File: "t2"}}, {{File: "q2"}}}
	s := nextSections(old, sp, []bool{true, false}, out)
	if len(s) != 2 || s[0].Name != "TRIPLE" || s[0].Tag != 0x100A || len(s[0].Runs) != 2 || len(s[0].Hash) != 64 || s[0].Appends == nil {
		t.Fatal(s[0])
	}
	if s[1].Runs[0].File != "q" || s[1].Appends[0].File != "q2" || s[1].Dead != 3 {
		t.Fatal(s[1])
	}
}
