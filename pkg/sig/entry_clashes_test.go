//ff:func feature=sig type=scanner control=sequence
//ff:what TestEntryClashes — same code, a group over a leaf, overlapping ranges
package sig

import "testing"

func TestEntryClashes(t *testing.T) {
	a := &Entry{Code: 1, Len: 2, Values: [][]byte{{1}}}
	if a.clashes(&Entry{Code: 2, Len: 2}, 4) {
		t.Fatal("different")
	}
	if !a.clashes(&Entry{Code: 1, Len: 2}, 4) {
		t.Fatal("same")
	}
	if !a.clashes(&Entry{Code: 4, Len: 4, Values: [][]byte{{2}}}, 4) {
		t.Fatal("both emit one code")
	}
	if a.clashes(&Entry{Code: 4, Len: 4}, 4) {
		t.Fatal("a group may share")
	}
	r := &Entry{Code: 4, Len: 4, Lo: 1, Hi: 5}
	if !r.clashes(&Entry{Code: 5, Len: 4, Lo: 4, Hi: 9}, 4) || r.clashes(&Entry{Code: 5, Len: 4, Lo: 5, Hi: 9}, 4) {
		t.Fatal("ranges")
	}
}
