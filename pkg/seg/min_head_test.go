//ff:func feature=seg type=helper control=sequence
//ff:what TestMinHead — the smallest non-nil row (first on ties), -1 when all are nil
package seg

import "testing"

func TestMinHead(t *testing.T) {
	if k := minHead([][]byte{nil, []byte("b"), []byte("a"), []byte("a")}); k != 2 {
		t.Fatal(k)
	}
	if k := minHead([][]byte{nil, nil}); k != -1 {
		t.Fatal(k)
	}
}
