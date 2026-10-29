//ff:func feature=seg type=encoder control=sequence
//ff:what TestAppendTable — key slot then start and count as big-endian u64, after a copy of the given bytes
package seg

import (
	"bytes"
	"testing"
)

func TestAppendTable(t *testing.T) {
	in := []byte{1}
	b := appendTable(in, []Part{{Key: []byte("k"), Start: 2, Count: 3}})
	want := []byte{1, 'k', 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 3}
	if !bytes.Equal(b, want) || len(in) != 1 {
		t.Fatal(b)
	}
}
