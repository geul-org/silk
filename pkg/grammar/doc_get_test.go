//ff:func feature=grammar type=scanner control=sequence
//ff:what TestDocGet — a path inside a LOCAL Object resolves through the document
package grammar

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocGet(t *testing.T) {
	g := testGrammar(t)
	text := RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab")))
	d, _ := Frame(edgeOf(0x00, tok("#1"), tok("#5"), text))
	l := edgeOf(0x04, tok("#3"), tok("Q1"), tok("#5"))
	if v, ok := d.Get(g, l, Path{2, 0}); !ok || string(v.B) != "ab" {
		t.Fatal(v, ok)
	}
	if v, ok := g.Get(l, Path{2, 0}); ok && string(v.B) == "ab" {
		t.Fatal("no document, no resolving: a LOCAL has only its ID")
	}
}
