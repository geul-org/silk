//ff:func feature=grammar type=scanner control=sequence
//ff:what TestGrammarCut — fields cut in order; a field before the target that does not fit stops the cut
package grammar

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestGrammarCut(t *testing.T) {
	g := testGrammar(t)
	tr := edgeOf(0x0A, tok("#1"), tok("P31"), tok("Q5"), tok("Q7"), []byte{0, 2})
	fs := g.Fields[0x100A]
	if v, ok := g.cut(tr, 0x100A, fs, 3, 6); !ok || !bytes.Equal(v.B, []byte{0, 2}) {
		t.Fatal(v, ok)
	}
	if _, ok := g.cut(tr, 0x100A, fs, 3, len(tr)+2); ok {
		t.Fatal("before the target")
	}
	if _, ok := g.cut(tr[:6], 0x100A, fs, 0, 6); ok {
		t.Fatal("at the target")
	}
	v, ok := g.cut(RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("abc"))), 0x0841, g.Fields[0x0841], 0, 2)
	if !ok || string(v.B) != "abc" {
		t.Fatal(v, ok)
	}
}
