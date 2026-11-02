//ff:func feature=sig type=scanner control=sequence
//ff:what TestDefCheckBits — no codes, past the end, shared bits; a flag needs none
package sig

import "testing"

func TestDefCheckBits(t *testing.T) {
	d := &Def{Bits: 4}
	if err := d.checkBits([]*Field{{Name: "a", Width: 2}}); err == nil {
		t.Fatal("no codes")
	}
	if err := d.checkBits([]*Field{{Name: "a", Width: 1, Kind: KindFlag}}); err != nil {
		t.Fatal(err)
	}
	if err := d.checkBits([]*Field{{Name: "a", Off: 3, Width: 2, Codes: []*Entry{{}}}}); err == nil {
		t.Fatal("past the end")
	}
}
