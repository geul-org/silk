//ff:func feature=cli type=scanner control=sequence
//ff:what TestLoadGrammar — from -grammar, from the archive's own set, and the failures
package cli

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestLoadGrammar(t *testing.T) {
	ctx := context.Background()
	f := newFlags("t", io.Discard)
	f.grammar = fixture.GrammarDir()
	if g, err := loadGrammar(ctx, f, "unused"); err != nil || len(g.Fields) == 0 {
		t.Fatal(err)
	}
	f.grammar = ""
	if _, err := loadGrammar(ctx, f, t.TempDir()); err == nil {
		t.Fatal("no archive")
	}
	data := fixture.Pack(t, fixture.Corpus(1, 1), nil)
	if _, err := loadGrammar(ctx, f, data); !errors.Is(err, grammar.ErrNoGrammar) {
		t.Fatal(err)
	}
	wk := filepath.Join(data, ".well-known", "geul", "0x0000")
	if err := os.CopyFS(wk, os.DirFS(fixture.GrammarDir())); err != nil {
		t.Fatal(err)
	}
	if g, err := loadGrammar(ctx, f, data); err != nil || len(g.Names) == 0 {
		t.Fatal(err)
	}
}
