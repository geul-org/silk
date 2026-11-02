//ff:func feature=index type=reader control=sequence
//ff:what Open — pin the current generation and map its files; a set generation the index does not cover is ErrStale; a cache of cacheBytes decoded documents
package index

import (
	"context"

	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/decl"
)

func Open(ctx context.Context, ar *archive.Archive, dir string, cacheBytes int64) (*Index, error) {
	g, pin, err := pinCurrent(dir, readCurrent)
	if err != nil {
		return nil, err
	}
	ix := &Index{Dir: dir, ar: ar, pin: pin, Decls: map[string]decl.Decl{}, secs: map[string][]*Sec{}, cache: newDocCache(cacheBytes)}
	if err := ix.load(g); err != nil {
		ix.Close()
		return nil, err
	}
	return ix, nil
}
