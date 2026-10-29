//ff:func feature=index type=encoder control=sequence
//ff:what TestEncodeBitmap — header and bits round-trip through readBitmap
package index

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestEncodeBitmap(t *testing.T) {
	p := filepath.Join(t.TempDir(), "l")
	in := Bitmap{B: []byte{0xFF, 0x03}, N: 10}
	os.WriteFile(p, encodeBitmap(in), 0o644)
	out, err := readBitmap(p)
	if err != nil || out.N != 10 || !bytes.Equal(out.B, in.B) {
		t.Fatal(err, out)
	}
}
