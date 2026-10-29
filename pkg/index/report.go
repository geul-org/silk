//ff:type feature=index type=model
//ff:what Report — what a build, update or compaction did: generation, documents (live, new, kept, removed), bytes read and written, rows, tasks, pause, time
package index

import "time"

// Report sums one writer run. NewBytes is the GEUL bytes of new document
// versions; Written counts every index byte written (segments, tables,
// bitmap, spec, manifest), so Written/NewBytes is the write amplification.
type Report struct {
	Generation uint64
	Unchanged  bool
	Docs       int
	New        int
	Kept       int
	Killed     int
	NewBytes   int64
	Read       int64
	Written    int64
	Rows       uint64
	Dropped    uint64
	Tasks      []string
	Paused     bool
	Elapsed    time.Duration
}
