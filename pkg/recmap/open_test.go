//ff:func feature=recmap type=reader control=sequence
//ff:what TestOpen — a written map opens; missing, unreadable and malformed files fail
package recmap

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	m, _, _ := sampleMap(t)
	dir := t.TempDir()
	p := filepath.Join(dir, "m.srm")
	WriteFile(p, m)
	f, err := Open(p)
	if err != nil || f.Header != m.Header {
		t.Fatal(err)
	}
	f.Close()
	if _, err := Open(filepath.Join(dir, "none")); err == nil {
		t.Fatal("missing")
	}
	os.Chmod(p, 0)
	if _, err := Open(p); err == nil && os.Geteuid() != 0 {
		t.Fatal("unreadable")
	}
	if _, err := Open(dir); err == nil {
		t.Fatal("a directory is not a map")
	}
}
