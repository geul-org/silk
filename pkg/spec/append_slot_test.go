//ff:func feature=spec type=encoder control=sequence
//ff:what TestAppendSlot — a short value is zero padded, a long one cut, after existing bytes
package spec

import (
	"bytes"
	"testing"
)

func TestAppendSlot(t *testing.T) {
	if b := AppendSlot([]byte{9}, []byte{1, 2}, 4); !bytes.Equal(b, []byte{9, 1, 2, 0, 0}) {
		t.Fatal(b)
	}
	if b := AppendSlot(nil, []byte{1, 2, 3, 4, 5, 6}, 4); !bytes.Equal(b, []byte{1, 2, 3, 4}) {
		t.Fatal(b)
	}
}
