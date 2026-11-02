//ff:func feature=sig type=scanner control=sequence
//ff:what TestPrefixGroups — a group without values may share a child's low code; a nested group asks for all of it
package sig

import "testing"

func TestPrefixGroups(t *testing.T) {
	d := testDef(t)
	f, _ := d.field("area")
	s, err := f.spans("asia")
	if err != nil || len(s) != 1 || s[0].Lo != 16 || s[0].Hi != 31 {
		t.Fatal(s, err)
	}
	if lo, hi := f.Codes[1].span(6); lo != 16 || hi != 19 {
		t.Fatal(lo, hi)
	}
}
