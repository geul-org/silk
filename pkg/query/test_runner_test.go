//ff:func feature=query type=builder control=sequence
//ff:what testRunner — a runner over an index
package query

import (
	"context"

	"github.com/geul-org/silk/pkg/index"
)

func testRunner(ix *index.Index) *runner {
	return &runner{ix: ix, ctx: context.Background(), seen: map[uint32]bool{}}
}
