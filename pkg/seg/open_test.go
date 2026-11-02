//ff:func feature=seg type=scanner control=sequence
//ff:what TestOpen — a missing, short, or inconsistent file fails; Map and Unmap on empty are harmless
package seg

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	dir := t.TempDir()
	if _, err := Open(filepath.Join(dir, "none")); err == nil {
		t.Fatal("missing")
	}
	short := filepath.Join(dir, "short")
	os.WriteFile(short, []byte("x"), 0o644)
	if _, err := Open(short); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	bad := filepath.Join(dir, "bad")
	os.WriteFile(bad, make([]byte, footerSize), 0o644)
	if _, err := Open(bad); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	empty := filepath.Join(dir, "empty")
	os.WriteFile(empty, nil, 0o644)
	if b, err := Map(empty); b != nil || err != nil {
		t.Fatal(b, err)
	}
	if err := Unmap(nil); err != nil {
		t.Fatal(err)
	}
	if b, err := Map(dir); err == nil {
		t.Fatal(b)
	}
}
