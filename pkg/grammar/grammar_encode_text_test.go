//ff:func feature=grammar type=scanner control=sequence
//ff:what TestGrammarEncodeText — each text form, and bad quoting or languages
package grammar

import (
	"errors"
	"testing"
)

func TestGrammarEncodeText(t *testing.T) {
	g := testGrammar(t)
	if v, err := g.encodeText(`"ab"`); err != nil || len(v) != 1 || Target(v[0]) != TagText {
		t.Fatal(v, err)
	}
	if v, err := g.encodeText(`"a"@en`); err != nil || len(v) != 2 || Target(v[0]) != TagLangOdd {
		t.Fatal(v, err)
	}
	if v, err := g.encodeText(`"a"@@en`); err != nil || len(v) != 1 {
		t.Fatal(v, err)
	}
	if _, err := g.encodeText(`"a`); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
	if _, err := g.encodeText(`"a"@@qqqqq`); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
}
