//ff:func feature=query type=encoder control=iteration dimension=1
//ff:what TestEncodeToken — Q, P and # Reference records, @tag and 0x records, * as nil; empty, unknown and non-numeric tokens fail
package query

import (
	"bytes"
	"encoding/hex"
	"errors"
	"testing"
)

func TestEncodeToken(t *testing.T) {
	want := map[string]string{"Q5": "00a00005", "P31": "00a8001f", "#7": "00880007", "Q169470": "01200002 95fe", "@de": "08c66465", "0x00a00005": "00a00005"}
	for tok, h := range want {
		w, _ := hex.DecodeString(stripSpaces(h))
		if b, err := encodeToken(tok); err != nil || !bytes.Equal(b, w) {
			t.Fatalf("%s: %x %v", tok, b, err)
		}
	}
	if b, err := encodeToken("*"); b != nil || err != nil {
		t.Fatal("wildcard")
	}
	for _, tok := range []string{"", "X1", "Qx", "@", "0xzz"} {
		if _, err := encodeToken(tok); !errors.Is(err, ErrSyntax) {
			t.Fatal(tok, err)
		}
	}
}
