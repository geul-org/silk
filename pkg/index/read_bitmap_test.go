//ff:func feature=index type=reader control=sequence
//ff:what TestReadBitmap — a written bitmap reads back; a missing file, a bad header or a body that does not match fails
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadBitmap(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "l")
	good := encodeBitmap(Bitmap{B: []byte{1}, N: 3})
	os.WriteFile(p, good, 0o644)
	if b, err := readBitmap(p); err != nil || !b.Get(0) {
		t.Fatal(err)
	}
	if _, err := readBitmap(filepath.Join(d, "none")); err == nil {
		t.Fatal("missing")
	}
	os.WriteFile(p, good[:20], 0o644)
	if _, err := readBitmap(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	bad := append([]byte{}, good...)
	bad[len(bad)-1] ^= 1
	os.WriteFile(p, bad, 0o644)
	if _, err := readBitmap(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
