//ff:func feature=seg type=scanner control=sequence
//ff:what TestRegionLen — rows by width
package seg

import "testing"

func TestRegionLen(t *testing.T) {
	if n := (Region{B: make([]byte, 12), W: 6}).Len(); n != 2 {
		t.Fatal(n)
	}
}
