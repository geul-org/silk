//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefField — common fields, block fields, nothing
package sig

import "testing"

func TestDefField(t *testing.T) {
	d := testDef(t)
	if f, w := d.field("type"); f == nil || w != nil {
		t.Fatal(f, w)
	}
	if f, w := d.field("pop"); f == nil || w == nil {
		t.Fatal(f, w)
	}
	if f, w := d.field("nope"); f != nil || w != nil {
		t.Fatal(f, w)
	}
}
