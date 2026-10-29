//ff:func feature=spec type=encoder control=sequence
//ff:what TestAppendKey — no path appends nothing; a path appends the field's slot
package spec

import (
	"bytes"
	"testing"
)

func TestAppendKey(t *testing.T) {
	tr := hexBytes("148a 0088 0005 00a8 001f 00a0 002a 00a0 0005 0000")
	if b := appendKey([]byte{7}, tr, Key{}); !bytes.Equal(b, []byte{7}) {
		t.Fatal(b)
	}
	if b := appendKey(nil, tr, Key{Path: []int{2}, Width: 2}); !bytes.Equal(b, hexBytes("00a0")) {
		t.Fatal(b)
	}
}
