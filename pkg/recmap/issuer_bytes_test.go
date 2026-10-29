//ff:func feature=recmap type=encoder control=iteration dimension=1
//ff:what TestIssuerBytes — 1/2/4/16-word Issuers decode; wrong width, prefix, case or digits fail
package recmap

import (
	"errors"
	"strings"
	"testing"
)

func TestIssuerBytes(t *testing.T) {
	if b, err := issuerBytes("0x00010000"); err != nil || len(b) != 4 || b[1] != 1 {
		t.Fatal(b, err)
	}
	if b, err := issuerBytes("0x" + strings.Repeat("0a", 32)); err != nil || len(b) != 32 {
		t.Fatal(b, err)
	}
	for _, s := range []string{"00010000", "0x000", "0x00AB", "0xzzzz"} {
		if _, err := issuerBytes(s); !errors.Is(err, ErrFormat) {
			t.Fatal(s, err)
		}
	}
}
