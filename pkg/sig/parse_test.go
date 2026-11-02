//ff:func feature=sig type=scanner control=sequence
//ff:what TestParse — offsets, codes and when blocks are read as written, a comment and blank line pass
package sig

import "testing"

func TestParse(t *testing.T) {
	d := testDef(t)
	if d.Name != "entity" || d.Bits != 64 || len(d.Common) != 2 || len(d.Whens) != 2 {
		t.Fatal(d)
	}
	f, w := d.field("born")
	if f.Off != 21 || f.Width != 4 || w == nil || w.Entry.Name != "person" || len(f.Codes) != 3 || f.Codes[1].Lo != 1800 || f.Codes[1].Hi != 1900 {
		t.Fatal(f, w)
	}
	if f, _ := d.field("area"); f.Codes[0].Code != 1 || f.Codes[0].Len != 2 || len(f.Codes[1].Values) != 1 {
		t.Fatal(f.Codes[0])
	}
	if f, w := d.field("none"); f != nil || w != nil {
		t.Fatal(f)
	}
	if d2, _ := Parse(defText + "\n"); d2.Hash == d.Hash {
		t.Fatal("hash of text")
	}
	if stripComment("a # b") != "a " {
		t.Fatal("comment")
	}
}
