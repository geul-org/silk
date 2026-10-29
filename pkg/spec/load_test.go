//ff:func feature=spec type=reader control=sequence
//ff:what TestLoad — a saved default spec loads equal; a missing file, unknown fields and an invalid spec fail
package spec

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "s.json")
	os.WriteFile(p, Default().Encode(), 0o644)
	s, err := Load(p)
	if err != nil || s.Hash() != Default().Hash() {
		t.Fatal(err)
	}
	if _, err := Load(filepath.Join(d, "none.json")); err == nil {
		t.Fatal("missing file")
	}
	os.WriteFile(p, []byte(`{"version":1,"tags":[],"x":1}`), 0o644)
	if _, err := Load(p); !errors.Is(err, ErrSpec) {
		t.Fatal(err)
	}
	os.WriteFile(p, []byte(`{"version":1,"tags":[]}`), 0o644)
	if _, err := Load(p); !errors.Is(err, ErrSpec) {
		t.Fatal(err)
	}
}
