//ff:func feature=index type=reader control=iteration dimension=1
//ff:what TestReadManifest — a committed manifest reads back; a missing one, bad JSON, another version or generation fail
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadManifest(t *testing.T) {
	d := t.TempDir()
	m := Manifest{Spec: "s"}
	commit(d, &m)
	if got, err := readManifest(d, 1); err != nil || got.Spec != "s" {
		t.Fatal(err)
	}
	if _, err := readManifest(d, 2); err == nil {
		t.Fatal("missing")
	}
	p := filepath.Join(d, manifestName(1))
	for _, b := range []string{"{", `{"version":1,"generation":1}`, `{"version":2,"generation":3}`} {
		os.WriteFile(p, []byte(b), 0o644)
		if _, err := readManifest(d, 1); !errors.Is(err, ErrFormat) {
			t.Fatal(b, err)
		}
	}
}
