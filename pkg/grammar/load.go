//ff:func feature=grammar type=reader control=iteration dimension=1
//ff:what Load — the grammar of the pack set directory of Issuer `0x0000`
package grammar

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func Load(ctx context.Context, dir string) (*Grammar, error) {
	a, err := pack.OpenDir(ctx, dir, pack.BuildOptions{})
	if err != nil {
		return nil, err
	}
	es, err := a.ReadDir("0x0000")
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %v", ErrNoGrammar, dir, err)
	}
	g := &Grammar{Fields: map[uint16][]Field{}, Names: map[string][]uint16{}, Odd: map[uint16]bool{}, Langs: map[string]uint16{}}
	d := resolve.Dir{Store: a}
	for _, e := range es {
		if err := g.addDoc(d, e.Name()); err != nil {
			return nil, err
		}
	}
	if len(g.Fields) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrNoGrammar, dir)
	}
	g.markOdd()
	return g, nil
}
