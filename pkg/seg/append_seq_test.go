//ff:func feature=seg type=encoder control=sequence
//ff:what TestAppendSeq — big-endian document then offset, so rows sort by them
package seg

import (
	"bytes"
	"testing"
)

func TestAppendSeq(t *testing.T) {
	if b := AppendSeq([]byte{9}, 0x01020304, 5); !bytes.Equal(b, []byte{9, 1, 2, 3, 4, 0, 0, 0, 5}) {
		t.Fatal(b)
	}
}
