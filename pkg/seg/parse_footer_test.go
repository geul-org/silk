//ff:func feature=seg type=scanner control=iteration dimension=1
//ff:what TestParseFooter — a footer round trips; damaged ones fail
package seg

import (
	"errors"
	"slices"
	"testing"
)

func TestParseFooter(t *testing.T) {
	b := encodeFooter(testInfo(), 2, 3, make([]byte, 32))
	i, p, r, err := parseFooter(b)
	if err != nil || i != testInfo() || p != 2 || r != 3 {
		t.Fatal(i, p, r, err)
	}
	for _, bad := range [][]byte{b[:9], append([]byte("XXXXXXXX"), b[8:]...), append(slices.Clone(b[:120]), "XXXXXXXX"...), append(append(slices.Clone(b[:8]), 0, 9), b[10:]...)} {
		if _, _, _, err := parseFooter(bad); !errors.Is(err, ErrFormat) {
			t.Fatal(err)
		}
	}
}
