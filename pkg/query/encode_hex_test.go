//ff:func feature=query type=encoder control=iteration dimension=1
//ff:what TestEncodeHex — one whole record passes; bad hex, an odd byte count, a truncated record or two records fail
package query

import (
	"errors"
	"testing"
)

func TestEncodeHex(t *testing.T) {
	if b, err := encodeHex("00a00005"); err != nil || len(b) != 4 {
		t.Fatal(err)
	}
	for _, h := range []string{"zz", "00a000", "00a0", "00a0000500a00005"} {
		if _, err := encodeHex(h); !errors.Is(err, ErrSyntax) {
			t.Fatal(h, err)
		}
	}
}
