//ff:func feature=recmap type=decoder control=sequence
//ff:what TestParseDocs — ascending documents parse; overlap and nonzero padding fail
package recmap

import (
	"errors"
	"testing"
)

func TestParseDocs(t *testing.T) {
	enc := func(ds ...DocEntry) []byte { return encodeBody(&Map{Docs: ds}) }
	if ds, err := parseDocs(enc(DocEntry{Off: 64, Len: 4, ID: 1}, DocEntry{Off: 68, Len: 2, ID: 2, Version: 3})); err != nil || len(ds) != 2 || ds[1].Version != 3 {
		t.Fatal(ds, err)
	}
	if _, err := parseDocs(enc(DocEntry{Off: 64, Len: 4}, DocEntry{Off: 66})); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	b := enc(DocEntry{Off: 64})
	b[20] = 1
	if _, err := parseDocs(b); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
