//ff:func feature=archive type=reader control=sequence
//ff:what Set.Record — a document's bytes by location, raw or zstd
package archive

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

func (s *Set) Record(ctx context.Context, l pack.Loc) ([]byte, error) {
	return s.r.Record(ctx, l)
}
