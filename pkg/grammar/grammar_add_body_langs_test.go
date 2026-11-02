//ff:func feature=grammar type=scanner control=sequence
//ff:what TestGrammarAddBodyLangs — the language registry of a body lands in Langs
package grammar

import "testing"

func TestGrammarAddBodyLangs(t *testing.T) {
	g := testGrammar(t)
	if len(g.Langs) < 1000 {
		t.Fatal(len(g.Langs))
	}
	b := Grammar{Fields: map[uint16][]Field{}, Names: map[string][]uint16{}, Langs: map[string]uint16{}}
	b.addBody(nil)
	if len(b.Fields) != 0 {
		t.Fatal(b.Fields)
	}
}
