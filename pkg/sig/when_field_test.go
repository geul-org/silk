//ff:func feature=sig type=scanner control=sequence
//ff:what TestWhenField — a block's field by name
package sig

import "testing"

func TestWhenField(t *testing.T) {
	w := &When{Fields: []*Field{{Name: "a"}}}
	if w.field("a") == nil || w.field("b") != nil {
		t.Fatal("field")
	}
}
