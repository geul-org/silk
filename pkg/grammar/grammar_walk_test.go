//ff:func feature=grammar type=scanner control=sequence
//ff:what TestGrammarWalk — steps, a raw stop, a missing field, resolving only below the top
package grammar

import (
	"bytes"
	"testing"
)

func TestGrammarWalk(t *testing.T) {
	g := testGrammar(t)
	r := testLabel()
	if v, ok := g.walk(nil, r, Path{2, 0}); !ok || !v.B2() {
		t.Fatal(v, ok)
	}
	if _, ok := g.walk(nil, r, Path{2, 1, 0}); ok {
		t.Fatal("raw")
	}
	if _, ok := g.walk(&Doc{}, r, Path{7}); ok {
		t.Fatal("missing")
	}
	if v, ok := g.walk(&Doc{}, r, nil); !ok || !bytes.Equal(v.B, r) {
		t.Fatal(v, ok)
	}
	if v, ok := g.walk(&Doc{refs: map[string][]byte{}}, r, Path{2, 0}); !ok || len(v.B) != 4 {
		t.Fatal(v, ok)
	}
}
