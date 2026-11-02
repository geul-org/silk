//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefCheck — placement of common and block fields
package sig

import "testing"

func TestDefCheck(t *testing.T) {
	on := &Field{Name: "a", Width: 4, Codes: []*Entry{{}}}
	d := &Def{Bits: 8, Common: []*Field{on}, Whens: []*When{{On: on}}}
	if err := d.check(); err != nil {
		t.Fatal(err)
	}
	d.Whens[0].Fields = []*Field{{Name: "b", Off: 2, Width: 4, Codes: []*Entry{{}}}}
	if err := d.check(); err == nil {
		t.Fatal("shares a bit")
	}
	d.Whens[0].Fields = nil
	d.Whens = append(d.Whens, &When{On: &Field{}})
	if err := d.check(); err == nil {
		t.Fatal("two fields")
	}
	d.Whens = nil
	d.Common = append(d.Common, &Field{Name: "c", Off: 1, Width: 1, Kind: KindFlag})
	if err := d.check(); err == nil {
		t.Fatal("common bit")
	}
}
