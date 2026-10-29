//ff:func feature=scan type=decoder control=iteration dimension=1
//ff:what TestNextField — a fitting field ends after its length; at the end, unframable or overlong fields fail
package scan

import "testing"

func TestNextField(t *testing.T) {
	rec := []byte{0x00, 0xa8, 0x00, 0x6a, 0x20, 0x00, 0x00, 0xa8}
	if e, ok := nextField(rec, 0); !ok || e != 4 {
		t.Fatal(e, ok)
	}
	for _, p := range []int{8, 4, 6} {
		if _, ok := nextField(rec, p); ok {
			t.Fatal(p)
		}
	}
}
