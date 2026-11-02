//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what TestRef — tokens as records; bad ones fail
package grammar

import (
	"bytes"
	"errors"
	"testing"
)

func TestRef(t *testing.T) {
	cases := map[string][]byte{"Q42": {0, 0xA0, 0, 42}, "P31": {0, 0xA8, 0, 31}, "#3": {0, 0x88, 0, 3}}
	for k, v := range cases {
		if b, err := Ref(k); err != nil || !bytes.Equal(b, v) {
			t.Fatal(k, b, err)
		}
	}
	for _, s := range []string{"Q", "X1", "Qx"} {
		if _, err := Ref(s); !errors.Is(err, ErrSyntax) {
			t.Fatal(s, err)
		}
	}
}
