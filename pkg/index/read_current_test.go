//ff:func feature=index type=reader control=sequence
//ff:what TestReadCurrent — the generation number; no CURRENT is ErrNoIndex, an unreadable one an error, garbage ErrFormat
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadCurrent(t *testing.T) {
	d := t.TempDir()
	if _, err := readCurrent(d); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(d, currentFile), []byte("12\n"), 0o644)
	if g, err := readCurrent(d); err != nil || g != 12 {
		t.Fatal(g, err)
	}
	os.WriteFile(filepath.Join(d, currentFile), []byte("x"), 0o644)
	if _, err := readCurrent(d); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	e := t.TempDir()
	os.Mkdir(filepath.Join(e, currentFile), 0o755)
	if _, err := readCurrent(e); err == nil || errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
}
