//ff:func feature=index type=scanner control=sequence
//ff:what TestManifest — commit, CURRENT and manifest reading, with damaged files refused
package index

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestManifest(t *testing.T) {
	dir := t.TempDir()
	if _, err := readCurrent(dir); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	if m, err := currentManifest(dir); err != nil || m.Paths.Name != "paths" || m.DocCount() != 0 {
		t.Fatal(m, err)
	}
	m := Manifest{Docs: []DocFile{{First: 3, Count: 4}}}
	if err := commit(dir, &m); err != nil || m.Generation != 1 || m.DocCount() != 7 {
		t.Fatal(err, m)
	}
	if g, err := readCurrent(dir); g != 1 || err != nil {
		t.Fatal(g, err)
	}
	if got, err := readManifest(dir, 1); err != nil || got.DocCount() != 7 {
		t.Fatal(got, err)
	}
	if _, err := readManifest(dir, 9); err == nil {
		t.Fatal("missing")
	}
	if n := m.nextName(dirSeg, "sseg"); !strings.HasPrefix(n, "seg/") || m.NextFile != 1 {
		t.Fatal(n)
	}
	os.WriteFile(filepath.Join(dir, manifestName(2)), []byte("{"), 0o644)
	if _, err := readManifest(dir, 2); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(dir, manifestName(3)), []byte(`{"version":9}`), 0o644)
	if _, err := readManifest(dir, 3); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(dir, currentFile), []byte("x"), 0o644)
	if _, err := readCurrent(dir); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if _, err := currentManifest(dir); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(dir, currentFile), []byte("9"), 0o644)
	if _, err := currentManifest(dir); err == nil {
		t.Fatal("missing generation")
	}
	os.Remove(filepath.Join(dir, currentFile))
	os.Mkdir(filepath.Join(dir, currentFile), 0o755)
	if _, err := readCurrent(dir); err == nil {
		t.Fatal("CURRENT is a directory")
	}
}
