//ff:func feature=index type=reader control=sequence
//ff:what Index.Bytes — a document's bytes, from the cache or its pack
package index

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
)

func (ix *Index) Bytes(ctx context.Context, seq uint32) ([]byte, error) {
	if b, ok := ix.cache.get(seq); ok {
		return b, nil
	}
	e := ix.Docs.At(seq)
	p := ix.Manifest.Packs[e.Pack]
	s := ix.ar.Set(p.Set)
	obj, ok := s.ObjectOf(p.Hash)
	if !ok {
		return nil, fmt.Errorf("%w: pack %s", ErrStale, p.Hash)
	}
	b, err := s.Record(ctx, pack.Loc{Obj: obj, Off: e.Off, Len: e.Len, Enc: e.Enc})
	if err != nil {
		return nil, err
	}
	ix.cache.put(seq, b)
	return b, nil
}
