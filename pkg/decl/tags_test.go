//ff:func feature=decl type=scanner control=iteration dimension=1
//ff:what TestTags — names, words, and what is neither
package decl

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestTags(t *testing.T) {
	g := fixture.Grammar(t)
	if ts, err := Tags(g, "UTF8-TEXT"); err != nil || len(ts) != 2 {
		t.Fatal(ts, err)
	}
	if ts, err := Tags(g, "0x1004"); err != nil || ts[0] != 0x1004 {
		t.Fatal(ts, err)
	}
	for _, s := range []string{"NOPE", "0x12", "0xzzzz", "0x7777"} {
		if _, err := Tags(g, s); !errors.Is(err, ErrDecl) {
			t.Fatal(s, err)
		}
	}
}
