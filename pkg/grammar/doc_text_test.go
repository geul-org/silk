//ff:func feature=grammar type=scanner control=sequence
//ff:what TestDocText — language text, plain text, and a raw or missing Object
package grammar

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocText(t *testing.T) {
	g := testGrammar(t)
	d := &Doc{}
	l, x, ok := d.Text(g, testLabel())
	if !ok || string(x) != "Hello" || !bytes.Equal(l, []byte{0x08, 0xC6, 'e', 'n'}) {
		t.Fatal(l, x, ok)
	}
	b := edgeOf(0x03, tok("#1"), tok("Q1"), RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("abc"))))
	if l, x, ok := d.Text(g, b); !ok || l != nil || string(x) != "abc" {
		t.Fatal(l, x, ok)
	}
	if _, _, ok := d.Text(g, edgeOf(0x04, tok("#1"), tok("Q1"))); ok {
		t.Fatal("no Object")
	}
	if _, _, ok := d.Text(g, edgeOf(0x0A, tok("#1"), tok("P1"), tok("Q1"), tok("Q2"), []byte{0, 0})); ok {
		t.Fatal("no text at field 2")
	}
}
