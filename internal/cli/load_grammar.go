//ff:func feature=cli type=reader control=sequence
//ff:what loadGrammar — the grammar of -grammar, or of the archive's `0x0000` set
package cli

import (
	"context"
	"fmt"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/grammar"
)

func loadGrammar(ctx context.Context, f *flags, dir string) (*grammar.Grammar, error) {
	if f.grammar != "" {
		return grammar.Load(ctx, f.grammar)
	}
	a, err := archive.Open(ctx, dir)
	if err != nil {
		return nil, err
	}
	s := a.Set("0x0000")
	if s == nil {
		return nil, fmt.Errorf("%w: the archive has no 0x0000 set (use -grammar DIR)", grammar.ErrNoGrammar)
	}
	return grammar.Load(ctx, s.Dir)
}
