//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what EdgeID — the Edge Id number of a hit's record; false without a record or a LOCAL Edge Id
package query

import (
	"context"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/index"
)

func EdgeID(ctx context.Context, ix *index.Index, h Hit) (uint64, bool, error) {
	rec, err := Record(ctx, ix, h)
	if err != nil || rec == nil {
		return 0, false, err
	}
	v, ok := ix.Grammar.Get(rec, grammar.Path{0})
	if !ok {
		return 0, false, nil
	}
	var n uint64
	for _, b := range v.B[2:] {
		n = n<<8 | uint64(b)
	}
	return n, true, nil
}
