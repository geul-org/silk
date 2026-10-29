//ff:func feature=recmap type=helper control=sequence
//ff:what TestGrow — a small buffer is replaced, a large one resliced in place
package recmap

import "testing"

func TestGrow(t *testing.T) {
	if b := grow(make([]byte, 2), 5); len(b) != 5 {
		t.Fatal(len(b))
	}
	big := make([]byte, 10)
	if b := grow(big, 3); len(b) != 3 || &b[0] != &big[0] {
		t.Fatal(len(b))
	}
}
