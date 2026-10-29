//ff:func feature=recmap type=reader control=sequence
//ff:what TestRead — an encoded map reads back; header read, header, size, table read, key table and document table failures surface
package recmap

import (
	"bytes"
	"errors"
	"testing"
)

func TestRead(t *testing.T) {
	m, _, _ := sampleMap(t)
	b, _ := m.Encode()
	n := int64(len(b))
	if f, err := Read(bytes.NewReader(b), n); err != nil || len(f.Docs) != 4 {
		t.Fatal(err)
	}
	if _, err := Read(bytes.NewReader(b[:10]), 10); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	bad := append([]byte("X"), b[1:]...)
	if _, err := Read(bytes.NewReader(bad), n); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if _, err := Read(bytes.NewReader(b), n-1); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if _, err := Read(failReader{b: b, from: headerSize}, n); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	k := &Map{Header: m.Header, Keys: []KeyRange{{Key: 1, Count: 2}}, Offs: []uint32{1}}
	kb, _ := k.Encode()
	if _, err := Read(bytes.NewReader(kb), int64(len(kb))); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	d := &Map{Header: m.Header, Docs: []DocEntry{{Off: 9, Len: 1}, {Off: 2}}}
	db, _ := d.Encode()
	if _, err := Read(bytes.NewReader(db), int64(len(db))); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
