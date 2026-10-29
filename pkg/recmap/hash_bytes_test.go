//ff:func feature=recmap type=encoder control=iteration dimension=1
//ff:what TestHashBytes — 64 lowercase hex digits decode; short, uppercase and non-hex fail
package recmap

import (
	"errors"
	"strings"
	"testing"
)

func TestHashBytes(t *testing.T) {
	if h, err := hashBytes(strings.Repeat("0f", 32)); err != nil || h[31] != 0x0f {
		t.Fatal(h, err)
	}
	for _, s := range []string{"0f", strings.Repeat("0F", 32), strings.Repeat("zz", 32)} {
		if _, err := hashBytes(s); !errors.Is(err, ErrFormat) {
			t.Fatal(s, err)
		}
	}
}
