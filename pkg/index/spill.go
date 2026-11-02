//ff:type feature=index type=model
//ff:what spill — one section's rows on their way to a sorted run: a memory buffer and the runs already written
package index

import (
	"sync/atomic"

	"github.com/geul-org/silk/pkg/seg"
)

type spill struct {
	info seg.Info
	tmp  string
	buf  *seg.Buffer
	runs []string
	n    *atomic.Uint64
}
