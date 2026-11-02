//ff:func feature=index type=builder control=sequence
//ff:what newSpill — an empty spill writing runs under tmp
package index

import (
	"sync/atomic"

	"github.com/geul-org/silk/pkg/seg"
)

func newSpill(info seg.Info, tmp string, n *atomic.Uint64) *spill {
	return &spill{info: info, tmp: tmp, buf: seg.NewBuffer(info), n: n}
}
