//ff:func feature=query type=helper control=sequence
//ff:what craftPart — an index part over hand-made pack bytes: one TRIPLE at offset 0, with or without a document entry, map read through wrap
package query

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/recmap"
)

// craftPart writes data to a temp file and pairs it with a map whose TRIPLE
// range is [0]. wrap may replace the map's reader (nil keeps bytes.Reader).
func craftPart(t *testing.T, data []byte, withDoc bool, wrap func([]byte) io.ReaderAt) index.Part {
	t.Helper()
	m := &recmap.Map{Header: recmap.Header{Issuer: "0x0001", Manifest: strings.Repeat("a", 64), Pack: strings.Repeat("b", 64), PackSize: uint64(len(data))},
		Keys: []recmap.KeyRange{{Key: tripleKey, Count: 1}}, Offs: []uint32{0}}
	if withDoc {
		m.Docs = []recmap.DocEntry{{Off: 0, Len: uint32(len(data)), ID: 9, Version: 1}}
	}
	b, err := m.Encode()
	if err != nil {
		t.Fatal(err)
	}
	var r io.ReaderAt = bytes.NewReader(b)
	if wrap != nil {
		r = wrap(b)
	}
	f, err := recmap.Read(r, int64(len(b)))
	if err != nil {
		t.Fatal(err)
	}
	p := filepath.Join(t.TempDir(), "pack")
	os.WriteFile(p, data, 0o644)
	d, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { d.Close() })
	return index.Part{Map: f, Data: d}
}
