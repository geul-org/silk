//ff:func feature=recmap type=decoder control=iteration dimension=1
//ff:what TestParseHeader — a written header parses; short, wrong magic or version, bad Issuer length and nonzero padding fail
package recmap

import (
	"errors"
	"strings"
	"testing"
)

func TestParseHeader(t *testing.T) {
	h := Header{Issuer: "0x0001", Manifest: strings.Repeat("ab", 32), Pack: strings.Repeat("cd", 32)}
	b, _ := encodeHeader(h)
	if got, err := parseHeader(b); err != nil || got != h {
		t.Fatal(got, err)
	}
	mut := func(i int, v byte) []byte { c := append([]byte(nil), b...); c[i] = v; return c }
	for _, bad := range [][]byte{b[:100], mut(0, 'X'), mut(8, 9), mut(10, 3), mut(120, 1)} {
		if _, err := parseHeader(bad); !errors.Is(err, ErrFormat) {
			t.Fatal(err)
		}
	}
}
