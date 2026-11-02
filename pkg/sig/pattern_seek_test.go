//ff:func feature=sig type=scanner control=sequence
//ff:what TestPatternSeek — next allowed codes, carries into free bits, and the end of the bits
package sig

import "testing"

func TestPatternSeek(t *testing.T) {
	p := Pattern{Bits: 8, Terms: []Term{{Off: 4, Width: 4, Spans: []Span{{2, 3}}}}}
	b := []byte{0x15}
	if !p.seek(b) || b[0] != 0x22 {
		t.Fatalf("%x", b)
	}
	b = []byte{0xF5}
	if p.seek(b) {
		t.Fatal("overflow")
	}
	b = []byte{0x23}
	if !p.seek(b) || b[0] != 0x23 {
		t.Fatal("already there")
	}
	if n, ok := (Term{Off: 0, Width: 4, Spans: []Span{{2, 3}}}).settle([]byte{0x40}, 3); ok || n != 0 {
		t.Fatal(n, ok)
	}
	if v, ok := nextCode([]Span{{5, 6}, {2, 3}}, 4); !ok || v != 5 {
		t.Fatal(v, ok)
	}
	if _, ok := nextCode(nil, 4); ok {
		t.Fatal("none")
	}
	c := []byte{0xFF}
	clearFrom(c, 4)
	if c[0] != 0xF0 {
		t.Fatal(c)
	}
	if carry(c, 4) || c[0] != 0 {
		t.Fatal("overflow", c)
	}
	if d := []byte{0x10}; !carry(d, 4) || d[0] != 0x20 {
		t.Fatal(d)
	}
	if allows([]Span{{1, 2}}, 3) || !allows([]Span{{1, 2}}, 2) {
		t.Fatal("allows")
	}
}
