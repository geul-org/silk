//ff:func feature=archive type=reader control=sequence
//ff:what Open — the pack archive of a data folder or `.well-known/geul/` directory, every set at its latest generation
package archive

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
)

func Open(ctx context.Context, dir string) (*Archive, error) {
	wk := wellKnown(dir)
	if !pack.IsPacked(wk) {
		return nil, fmt.Errorf("%w: %s", ErrNotPacked, dir)
	}
	a, err := pack.OpenDir(ctx, wk, pack.BuildOptions{})
	if err != nil {
		return nil, err
	}
	sets, err := openSets(ctx, a.Sets())
	return &Archive{Root: wk, Sets: sets}, err
}
