//ff:func feature=recmap type=encoder control=iteration dimension=1
//ff:what TestEncodeHeader — 176 bytes that parse back; bad Issuer, pack or manifest hex fail
package recmap

import (
	"errors"
	"strings"
	"testing"
)

func TestEncodeHeader(t *testing.T) {
	h := Header{Issuer: "0x00010000", Generation: 3, Manifest: strings.Repeat("ab", 32), Pack: strings.Repeat("cd", 32), PackSize: 99, Keys: 1, Records: 2, Docs: 3}
	h.Sum[0] = 1
	b, err := encodeHeader(h)
	if err != nil || len(b) != headerSize {
		t.Fatal(len(b), err)
	}
	if got, err := parseHeader(b); err != nil || got != h {
		t.Fatal(got, err)
	}
	for _, bad := range []Header{{Issuer: "0x1"}, {Issuer: "0x0001", Pack: "x"}, {Issuer: "0x0001", Pack: h.Pack, Manifest: "x"}} {
		if _, err := encodeHeader(bad); !errors.Is(err, ErrFormat) {
			t.Fatal(bad, err)
		}
	}
}
