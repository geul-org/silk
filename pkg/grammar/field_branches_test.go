//ff:func feature=grammar type=scanner control=sequence
//ff:what TestFieldBranches — no layout, a broken Edge Id, Edge Id itself, a field past the layout, a Value's field
package grammar

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFieldBranches(t *testing.T) {
	g := testGrammar(t)
	if _, ok := g.field([]byte{0x1F, 0x7F}, 0); ok {
		t.Fatal("no layout")
	}
	if _, ok := g.field([]byte{0x10, 0x84}, 0); ok {
		t.Fatal("no Edge Id")
	}
	if _, ok := g.field([]byte{0x11, 0x04, 0xFF, 0xFF}, 0); ok {
		t.Fatal("broken Edge Id")
	}
	r := testLabel()
	if v, ok := g.field(r, 0); !ok || !bytes.Equal(v.B, tok("#3")) {
		t.Fatal(v, ok)
	}
	if v, ok := g.field(r, 1); !ok || !bytes.Equal(v.B, tok("Q42")) {
		t.Fatal(v, ok)
	}
	if _, ok := g.field(r, 3); ok {
		t.Fatal("past the fields")
	}
	if _, ok := g.field(RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab"))), -1); ok {
		t.Fatal("negative")
	}
}
