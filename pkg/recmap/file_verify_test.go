//ff:func feature=recmap type=rule control=sequence
//ff:what TestFileVerify — a fresh map verifies; body read, body hash, rebuild, off[] read and table mismatches fail
package recmap

import (
	"bytes"
	"errors"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestFileVerify(t *testing.T) {
	m, p, data := sampleMap(t)
	b, _ := m.Encode()
	read := func(r interface {
		ReadAt([]byte, int64) (int, error)
	}) *File {
		f, err := Read(bytes.NewReader(b), int64(len(b)))
		if err != nil {
			t.Fatal(err)
		}
		f.r = r
		return f
	}
	if err := read(bytes.NewReader(b)).Verify(data, p.Docs); err != nil {
		t.Fatal(err)
	}
	if err := read(failReader{b: b, from: headerSize}).Verify(data, p.Docs); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	flip := append([]byte(nil), b...)
	flip[len(flip)-1] ^= 1
	if err := read(bytes.NewReader(flip)).Verify(data, p.Docs); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := read(bytes.NewReader(b)).Verify(data, []archive.Doc{p.Docs[1], p.Docs[0]}); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := read(offsFail{b: b, at: headerSize + int64(m.Header.Keys)*keySize + int64(m.Header.Docs)*docSize}).Verify(data, p.Docs); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := read(bytes.NewReader(b)).Verify(data, p.Docs[:3]); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
}
