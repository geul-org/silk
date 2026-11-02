//ff:func feature=fixture type=scanner control=sequence
//ff:what TestLangText — language then text in one Value
package fixture

import (
	"bytes"
	"testing"
)

func TestLangText(t *testing.T) {
	b := LangText("en", "x")
	if !bytes.Contains(b, []byte("en")) || !bytes.Contains(b, []byte("x")) {
		t.Fatal(b)
	}
}
