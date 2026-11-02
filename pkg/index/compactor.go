//ff:type feature=index type=model
//ff:what compactor — one compaction run
package index

import "sync/atomic"

type compactor struct {
	dir  string
	o    Options
	m    Manifest
	live Bitmap
	full bool
	tmp  string
	n    atomic.Uint64
}
