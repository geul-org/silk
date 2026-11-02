//ff:func feature=sig type=scanner control=sequence
//ff:what TestFieldClash — a name or a code taken
package sig

import "testing"

func TestFieldClash(t *testing.T) {
	f := &Field{Width: 4, Codes: []*Entry{{Name: "a", Code: 4, Len: 4}}}
	if f.clash(&Entry{Name: "b", Code: 5, Len: 4}) != nil {
		t.Fatal("free")
	}
	if f.clash(&Entry{Name: "a", Code: 5, Len: 4}) == nil || f.clash(&Entry{Name: "b", Code: 4, Len: 4}) == nil {
		t.Fatal("taken")
	}
}
