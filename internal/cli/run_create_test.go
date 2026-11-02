//ff:func feature=cli type=scanner control=iteration dimension=1
//ff:what TestRunCreate — a declared section reaches the manifest; each refusal comes through
package cli

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/index"
)

func TestRunCreate(t *testing.T) {
	e, _ := testEnv("")
	data, ix := project(t)
	args := append([]string{"text", "alias", "ALIAS", "-grammar", fixture.GrammarDir()}, append(ix, data)...)
	if err := runCreate(e, args); err != nil {
		t.Fatal(err)
	}
	if err := runCreate(e, args); !errors.Is(err, index.ErrExists) {
		t.Fatal(err)
	}
	for _, a := range [][]string{{"-nope"}, {"text"}, {"text", "x", "NOPE", "-grammar", fixture.GrammarDir(), data}, {"text", "x", "LABEL", t.TempDir()}} {
		if err := runCreate(e, a); err == nil {
			t.Fatal(a)
		}
	}
}
