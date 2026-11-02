//ff:func feature=fixture type=rule control=sequence
//ff:what TestGrammarAt — a directory without a grammar fails the test
package fixture

import "testing"

func TestGrammarAt(t *testing.T) {
	if !fails(t, func(tb testing.TB) { grammarAt(tb, t.TempDir()) }) {
		t.Fatal("no grammar")
	}
	if g := Grammar(t); len(g.Fields) == 0 {
		t.Fatal("empty")
	}
}
