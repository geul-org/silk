//ff:func feature=sig type=scanner control=sequence
//ff:what TestFieldEntry — a code by name
package sig

import "testing"

func TestFieldEntry(t *testing.T) {
	f := &Field{Codes: []*Entry{{Name: "a"}}}
	if f.entry("a") == nil || f.entry("b") != nil {
		t.Fatal("entry")
	}
}
