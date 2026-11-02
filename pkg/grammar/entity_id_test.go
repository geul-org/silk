//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what TestEntityID — QID and PID fit; LOCAL, variable, wide and short do not
package grammar

import "testing"

func TestEntityID(t *testing.T) {
	if n := EntityID(tok("Q42")); n != 0x4000002A {
		t.Fatalf("%x", n)
	}
	if n := EntityID(tok("P31")); n != 0x5000001F {
		t.Fatalf("%x", n)
	}
	if n := EntityID(tok("Q141796311")); n != 0x40000000|141796311 {
		t.Fatalf("%x", n)
	}
	for _, r := range [][]byte{tok("#1"), tok("Q268435456"), {0, 0xA4, 0, 1}, {0, 1}, {0, 0xA0, 0, 1, 0, 0, 0, 0}, {0x0A, 0xC0, 0, 1}} {
		if n := EntityID(r); n != 0 {
			t.Fatal(r, n)
		}
	}
}
