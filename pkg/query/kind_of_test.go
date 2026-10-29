//ff:func feature=query type=helper control=sequence
//ff:what TestKindOf — Edge kind and codec of a layout name; false for an unknown name
package query

import "testing"

func TestKindOf(t *testing.T) {
	if k, c, ok := kindOf("SOURCE"); !ok || k != 2 || c != 0x20 {
		t.Fatal(k, c, ok)
	}
	if _, _, ok := kindOf("BIND"); ok {
		t.Fatal("BIND has no layout")
	}
}
