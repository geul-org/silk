//ff:func feature=query type=encoder control=iteration dimension=1
//ff:what TestEncodeToken — Q, P and # Reference records in minimal words; * is nil; empty, unknown prefixes and non-numbers fail
package query

import (
	"bytes"
	"errors"
	"testing"
)

func TestEncodeToken(t *testing.T) {
	cases := map[string][]byte{
		"Q42": {0x00, 0xa0, 0x00, 0x2a},
		"P31": {0x00, 0xa8, 0x00, 0x1f},
		"#7":  {0x00, 0x88, 0x00, 0x07},
		"*":   nil,
	}
	for tok, want := range cases {
		if got, err := encodeToken(tok); err != nil || !bytes.Equal(got, want) {
			t.Fatalf("%s %x %v", tok, got, err)
		}
	}
	for _, tok := range []string{"", "X1", "Qx"} {
		if _, err := encodeToken(tok); !errors.Is(err, ErrSyntax) {
			t.Fatal(tok, err)
		}
	}
}
