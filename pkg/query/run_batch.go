//ff:func feature=query type=builder control=iteration dimension=1
//ff:what RunBatch — answer each query in turn; the first error stops
package query

import (
	"context"

	"github.com/geul-org/silk/pkg/index"
)

func RunBatch(ctx context.Context, ix *index.Index, qs []Query) ([]Result, error) {
	out := make([]Result, 0, len(qs))
	for _, q := range qs {
		r, err := Run(ctx, ix, q)
		if err != nil {
			return out, err
		}
		out = append(out, r)
	}
	return out, nil
}
