//ff:func feature=index type=encoder control=sequence
//ff:what TestAppendPairs — 6 bytes per (TAG, count), big-endian
package index

import (
	"bytes"
	"testing"
)

func TestAppendPairs(t *testing.T) {
	if b := appendPairs(nil, []TagCount{{0x100A, 3}}); !bytes.Equal(b, []byte{0x10, 0x0A, 0, 0, 0, 3}) {
		t.Fatal(b)
	}
}
