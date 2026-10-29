//ff:func feature=scan type=reader control=sequence
//ff:what TestWindowRecord — frames records across window refills; bad TAG, short pack and read errors fail
package scan

import (
	"bytes"
	"errors"
	"testing"
)

func TestWindowRecord(t *testing.T) {
	rec, _, _ := tripleRec(7, 106, 42, 169470)
	data := append(append([]byte{0x08, 0x82, 0x00, 0x05}, rec...), 0x18, 0x0F)
	w := &window{data: bytes.NewReader(data), size: uint64(len(data)), min: 4}
	if b, err := w.record(0); err != nil || len(b) != 4 {
		t.Fatal(b, err)
	}
	if b, err := w.record(4); err != nil || !bytes.Equal(b, rec) {
		t.Fatal(b, err)
	}
	if b, err := w.record(uint64(len(data) - 2)); err != nil || len(b) != 2 {
		t.Fatal(b, err)
	}
	if _, err := w.record(uint64(len(data))); !errors.Is(err, ErrRecord) {
		t.Fatal(err)
	}
	bad := &window{data: bytes.NewReader([]byte{0xFF, 0xFF}), size: 2, min: 4}
	if _, err := bad.record(0); !errors.Is(err, ErrRecord) {
		t.Fatal(err)
	}
	cut := &window{data: bytes.NewReader(rec[:8]), size: 8, min: 4}
	if _, err := cut.record(0); !errors.Is(err, ErrRecord) {
		t.Fatal(err)
	}
	flaky := &window{data: &failAfter{r: bytes.NewReader(data), n: 1}, size: uint64(len(data)), min: 2}
	if _, err := flaky.record(4); err == nil {
		t.Fatal("second read must fail")
	}
}
