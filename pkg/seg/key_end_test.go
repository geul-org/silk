//ff:func feature=seg type=helper control=sequence
//ff:what TestKeyEnd — the end of a run of equal keys
package seg

import "testing"

func TestKeyEnd(t *testing.T) {
	gs := []Group{{Key: []byte("a")}, {Key: []byte("a")}, {Key: []byte("b")}}
	if keyEnd(gs, 0) != 2 || keyEnd(gs, 2) != 3 {
		t.Fatal(keyEnd(gs, 0), keyEnd(gs, 2))
	}
}
