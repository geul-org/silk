//ff:func feature=query type=encoder control=sequence
//ff:what TestEncodeLang — even and odd BCP 47 tags as LANGUAGE-BCP47-CODE records; an empty tag fails
package query

import (
	"bytes"
	"errors"
	"testing"
)

func TestEncodeLang(t *testing.T) {
	if b, _ := encodeLang("ko"); !bytes.Equal(b, []byte{0x08, 0xc6, 'k', 'o'}) {
		t.Fatalf("%x", b)
	}
	if b, _ := encodeLang("ast"); !bytes.Equal(b, []byte{0x09, 0x47, 'a', 's', 't', 0}) {
		t.Fatalf("%x", b)
	}
	if _, err := encodeLang(""); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
}
