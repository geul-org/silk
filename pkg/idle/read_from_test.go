//ff:func feature=idle type=reader control=sequence
//ff:what TestReadFrom — a Sample from given files; missing files and a bad loadavg fail
package idle

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadFrom(t *testing.T) {
	d := t.TempDir()
	la, mi := filepath.Join(d, "loadavg"), filepath.Join(d, "meminfo")
	os.WriteFile(la, []byte("2.5 1 1 1/2 3\n"), 0o644)
	os.WriteFile(mi, []byte("MemTotal: 4 kB\nMemAvailable: 2 kB\n"), 0o644)
	s, err := readFrom(la, mi, 4)
	if err != nil || s != (Sample{Load1: 2.5, CPUs: 4, MemTotal: 4096, MemAvail: 2048}) {
		t.Fatal(s, err)
	}
	if _, err := readFrom(filepath.Join(d, "x"), mi, 4); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	if _, err := readFrom(la, filepath.Join(d, "x"), 4); !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	os.WriteFile(la, []byte(""), 0o644)
	if _, err := readFrom(la, mi, 4); err == nil {
		t.Fatal("bad loadavg")
	}
}
