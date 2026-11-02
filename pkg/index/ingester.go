//ff:type feature=index type=model
//ff:what ingester — the shared sink of worker rows, guarded by one lock
package index

import (
	"sync"
	"sync/atomic"
)

type ingester struct {
	u      *updater
	ts     []*target
	mu     sync.Mutex
	budget int64
	failed atomic.Bool
	err    error
	maxDoc uint64
	read   uint64
}
