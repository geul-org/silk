//ff:func feature=spec type=rule control=iteration dimension=1
//ff:what TestLayoutOf — every known Edge layout by codec, and none for another Kind or an unknown codec
package spec

import "testing"

func TestLayoutOf(t *testing.T) {
	want := map[uint8]string{0x00: "REFERENCE-NODE", 0x04: "LABEL", 0x05: "ALIAS", 0x06: "DESCRIPTION", 0x0A: "TRIPLE", 0x0B: "QUALIFIER", 0x20: "SOURCE"}
	for c, name := range want {
		l, ok := LayoutOf(KindEdge, c)
		if !ok || l.Name != name || len(l.Fields) < 3 {
			t.Fatalf("0x%02x: %v %v", c, l, ok)
		}
	}
	if _, ok := LayoutOf(KindEdge, 0x7F); ok {
		t.Fatal("unknown codec")
	}
	if _, ok := LayoutOf(1, 0x0A); ok {
		t.Fatal("Value kind")
	}
}
