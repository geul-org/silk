//ff:func feature=recmap type=reader control=sequence
//ff:what TestFileClose — a Read map closes to nil; an opened one closes its file once
package recmap

import (
	"path/filepath"
	"testing"
)

func TestFileClose(t *testing.T) {
	if err := (&File{}).Close(); err != nil {
		t.Fatal(err)
	}
	m, _, _ := sampleMap(t)
	p := filepath.Join(t.TempDir(), "m.srm")
	WriteFile(p, m)
	f, _ := Open(p)
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err == nil {
		t.Fatal("second close must fail")
	}
}
