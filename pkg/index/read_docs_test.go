//ff:func feature=index type=reader control=sequence
//ff:what TestReadDocs — a missing file, a bad header, a body hash mismatch and a bad entry fail
package index

import (
	"crypto/sha256"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadDocs(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "d")
	if _, _, err := readDocs(filepath.Join(d, "none")); err == nil {
		t.Fatal("missing")
	}
	good := encodeDocs(0, []DocEntry{{Len: 1, Counts: []TagCount{{1, 1}}}})
	os.WriteFile(p, good[:10], 0o644)
	if _, _, err := readDocs(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	bad := append([]byte{}, good...)
	bad[len(bad)-1] ^= 1
	os.WriteFile(p, bad, 0o644)
	if _, _, err := readDocs(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	ent := append([]byte{}, good...)
	ent[docsHead+38] = 1
	s := sha256.Sum256(ent[docsHead:])
	copy(ent[24:56], s[:])
	os.WriteFile(p, ent, 0o644)
	if _, _, err := readDocs(p); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
