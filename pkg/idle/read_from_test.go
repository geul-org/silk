//ff:func feature=idle type=scanner control=sequence
//ff:what TestReadFrom — files parse; missing or bad ones are ErrUnsupported; Read works on this machine or says so
package idle

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadFrom(t *testing.T) {
	dir := t.TempDir()
	l, m := filepath.Join(dir, "l"), filepath.Join(dir, "m")
	os.WriteFile(l, []byte("1.50 0 0"), 0o644)
	os.WriteFile(m, []byte("MemTotal: 10 kB\nMemAvailable: 5 kB\nx\nOther: 3 kB\n"), 0o644)
	s, err := readFrom(l, m, 2)
	if err != nil || s.Load1 != 1.5 || s.MemTotal != 10240 || s.MemAvail != 5120 || s.CPUs != 2 {
		t.Fatal(s, err)
	}
	if _, err := readFrom(filepath.Join(dir, "x"), m, 1); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	os.WriteFile(l, nil, 0o644)
	if _, err := readFrom(l, m, 1); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	os.WriteFile(l, []byte("x"), 0o644)
	if _, err := readFrom(l, m, 1); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	if _, err := Read(); err != nil && !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	if a, b := parseMeminfo(""); a != 0 || b != 0 {
		t.Fatal(a, b)
	}
}
