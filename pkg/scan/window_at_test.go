//ff:func feature=scan type=reader control=sequence
//ff:what TestWindowAt — reads, reuses the window, refills past it, caps at the pack end; beyond the end and read errors fail
package scan

import (
	"bytes"
	"errors"
	"testing"
)

func TestWindowAt(t *testing.T) {
	data := []byte("0123456789")
	w := &window{data: bytes.NewReader(data), size: 10, min: 4}
	if b, err := w.at(2, 2); err != nil || string(b) != "2345" {
		t.Fatal(string(b), err)
	}
	if b, err := w.at(3, 2); err != nil || string(b) != "345" || w.lo != 2 {
		t.Fatal(string(b), err)
	}
	if b, err := w.at(5, 6); err != nil || string(b) != "56789" {
		t.Fatal(string(b), err)
	}
	if b, err := w.at(6, 1); err != nil || string(b) != "6789" || w.lo != 5 {
		t.Fatal(string(b), err)
	}
	if _, err := w.at(10, 1); !errors.Is(err, ErrRecord) {
		t.Fatal(err)
	}
	w.data = errReader{}
	if _, err := w.at(0, 1); err == nil || len(w.buf) != 0 {
		t.Fatal(err)
	}
}
