//ff:func feature=grammar type=scanner control=sequence
//ff:what TestGrammarGet — fields of an Edge, of a Value inside it, and paths that do not exist
package grammar

import (
	"bytes"
	"testing"
)

func TestGrammarGet(t *testing.T) {
	g := testGrammar(t)
	r := testLabel()
	if v, ok := g.Get(r, Path{0}); !ok || !bytes.Equal(v.B, tok("#3")) {
		t.Fatal(v, ok)
	}
	if v, ok := g.Get(r, Path{1}); !ok || !bytes.Equal(v.B, tok("Q42")) {
		t.Fatal(v, ok)
	}
	if v, ok := g.Get(r, Path{2, 1}); !ok || string(v.B) != "Hello" || !v.Rest || !v.Raw {
		t.Fatal(v, ok)
	}
	if _, ok := g.Get(r, Path{2, 1, 0}); ok {
		t.Fatal("inside a raw field")
	}
	if _, ok := g.Get(r, Path{9}); ok {
		t.Fatal("past the fields")
	}
	if _, ok := g.Get([]byte{0x1F, 0x7F}, Path{0}); ok {
		t.Fatal("no layout")
	}
	short := []byte{0x10, 0x84}
	if _, ok := g.Get(short, Path{0}); ok {
		t.Fatal("no Edge Id")
	}
}
