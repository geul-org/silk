//ff:func feature=index type=reader control=sequence
//ff:what TestLoadDocs — the document tables in order; a missing file or a file out of sequence fails
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDocs(t *testing.T) {
	_, _, dir := built(t)
	m, _ := CurrentManifest(dir)
	if ds, err := loadDocs(dir, m); err != nil || len(ds) != 4 {
		t.Fatal(err)
	}
	m.Docs[0].First = 1
	if _, err := loadDocs(dir, m); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	os.Remove(filepath.Join(dir, m.Docs[0].File))
	if _, err := loadDocs(dir, m); err == nil {
		t.Fatal("missing")
	}
}
