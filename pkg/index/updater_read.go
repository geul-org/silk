//ff:func feature=index type=reader control=sequence
//ff:what updater.read — a document's bytes through its set's reader
package index

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
)

func (u *updater) read(ctx context.Context, e DocEntry) ([]byte, error) {
	p := u.m.Packs[e.Pack]
	s := u.ar.Set(p.Set)
	obj, ok := s.ObjectOf(p.Hash)
	if !ok {
		return nil, fmt.Errorf("%w: pack %s", ErrStale, p.Hash)
	}
	return s.Record(ctx, pack.Loc{Obj: obj, Off: e.Off, Len: e.Len, Enc: e.Enc})
}
