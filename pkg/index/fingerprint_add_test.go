//ff:func feature=index type=scanner control=sequence
//ff:what TestFingerprintAdd — order free, content bound
package index

import "testing"

func TestFingerprintAdd(t *testing.T) {
	var a, b, c fingerprint
	a.add([]byte("x"))
	a.add([]byte("y"))
	b.add([]byte("y"))
	b.add([]byte("x"))
	c.add([]byte("x"))
	c.add([]byte("z"))
	if a != b || a == c {
		t.Fatal(a, b, c)
	}
}
