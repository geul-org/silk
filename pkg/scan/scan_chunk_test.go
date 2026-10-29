//ff:func feature=scan type=scanner control=sequence
//ff:what TestScanChunk — matching records come back with copies of their bytes; a bad offset fails
package scan

import (
	"bytes"
	"errors"
	"testing"
)

func TestScanChunk(t *testing.T) {
	a, p, o := tripleRec(1, 106, 42, 169470)
	b, _, _ := tripleRec(2, 31, 42, 5)
	data := append(append([]byte(nil), a...), b...)
	w := &window{data: bytes.NewReader(data), size: uint64(len(data)), min: 8}
	hits, err := scanChunk(w, []uint32{0, uint32(len(a))}, Triple{Pred: p, Obj: o})
	if err != nil || len(hits) != 1 || hits[0].Off != 0 || !bytes.Equal(hits[0].Rec, a) {
		t.Fatal(hits, err)
	}
	if _, err := scanChunk(w, []uint32{uint32(len(data))}, Triple{}); !errors.Is(err, ErrRecord) {
		t.Fatal(err)
	}
}
