//ff:func feature=grammar type=scanner control=sequence
//ff:what TestDocResolve — a LOCAL follows its REFERENCE-NODE; another record and an unbound LOCAL stay
package grammar

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocResolve(t *testing.T) {
	g := testGrammar(t)
	text := RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab")))
	b := append(edgeOf(0x00, tok("#1"), tok("#5"), text), edgeOf(0x00, tok("#2"), tok("#6"))...)
	d, err := Frame(b)
	if err != nil {
		t.Fatal(err)
	}
	if o := d.Resolve(g, tok("#5")); !bytes.Equal(o, text) {
		t.Fatal(o)
	}
	if o := d.Resolve(g, tok("#9")); !bytes.Equal(o, tok("#9")) {
		t.Fatal(o)
	}
	if o := d.Resolve(g, tok("Q1")); !bytes.Equal(o, tok("Q1")) {
		t.Fatal(o)
	}
	l := edgeOf(0x04, tok("#3"), tok("Q1"), tok("#5"))
	if v, ok := d.Get(g, l, Path{2, 0}); !ok || string(v.B) != "ab" {
		t.Fatal(v, ok)
	}
	if _, _, ok := d.Text(g, edgeOf(0x04, tok("#3"), tok("Q1"), tok("#9"))); ok {
		t.Fatal("unresolved LOCAL is no text")
	}
}
