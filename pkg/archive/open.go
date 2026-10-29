//ff:func feature=archive type=reader control=iteration dimension=1
//ff:what Open — the pack archive under a data folder or `.well-known/geul/` directory, every set at its latest generation
package archive

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
)

// Open finds the pack sets of dir (a data folder holding `.well-known/geul/`,
// or that directory itself) through pack.OpenDir. A plain directory archive is
// ErrNotPacked.
func Open(ctx context.Context, dir string) (*Archive, error) {
	wk := wellKnown(dir)
	if !pack.IsPacked(wk) {
		return nil, fmt.Errorf("%w: %s", ErrNotPacked, dir)
	}
	a, err := pack.OpenDir(ctx, wk, pack.BuildOptions{})
	if err != nil {
		return nil, err
	}
	out := &Archive{Root: wk}
	for _, si := range a.Sets() {
		out.Sets = append(out.Sets, Set{Name: si.Name, Dir: si.Dir, Manifest: si.Manifest, Generation: si.Generation})
	}
	return out, nil
}
