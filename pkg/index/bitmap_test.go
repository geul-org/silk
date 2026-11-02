//ff:func feature=index type=scanner control=sequence
//ff:what TestBitmap — bits grow, clear, count and round trip; damaged files fail
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestBitmap(t *testing.T) {
	var b Bitmap
	b.Put(9, true)
	b.Put(3, true)
	b.Put(3, false)
	b.Put(100, false)
	if !b.Get(9) || b.Get(3) || b.Get(1000) || b.N != 101 || b.count() != 1 {
		t.Fatal(b)
	}
	dir := t.TempDir()
	p := filepath.Join(dir, "l")
	os.WriteFile(p, encodeBitmap(b), 0o644)
	r, err := readBitmap(p)
	if err != nil || !r.Get(9) || r.N != 101 {
		t.Fatal(r, err)
	}
	if _, err := readBitmap(filepath.Join(dir, "x")); err == nil {
		t.Fatal("missing")
	}
	os.WriteFile(p, []byte("short"), 0o644)
	if _, err := readBitmap(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	bad := encodeBitmap(b)
	bad[0] ^= 1
	os.WriteFile(p, bad, 0o644)
	if _, err := readBitmap(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	os.WriteFile(p, appendTail(nil, liveMagic, 0, 99, nil), 0o644)
	if _, err := readBitmap(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if _, _, _, err := readTail("x", appendTail(nil, "NOTMAGIC", 0, 0, nil), liveMagic); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
