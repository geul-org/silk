//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefWhenOf — a code with a block, without, on another field
package sig

import "testing"

func TestDefWhenOf(t *testing.T) {
	d := testDef(t)
	if w := d.whenOf(d.Common[0], 4); w == nil || w.Entry.Name != "person" {
		t.Fatal(w)
	}
	if w := d.whenOf(d.Common[0], 9); w != nil {
		t.Fatal(w)
	}
	if w := d.whenOf(d.Common[1], 4); w != nil {
		t.Fatal(w)
	}
}
