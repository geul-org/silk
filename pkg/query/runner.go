//ff:type feature=query type=model
//ff:what runner — one query's state: index, context and work counters
package query

import (
	"context"

	"github.com/geul-org/silk/pkg/index"
)

type runner struct {
	ix      *index.Index
	ctx     context.Context
	decoded int
	seen    map[uint32]bool
}
