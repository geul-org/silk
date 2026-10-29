//ff:type feature=index type=model
//ff:what ingest — documents being indexed: spec lookup table, one row bucket per tag (sorted or appended), segments written, bytes read and written, open data packs, pack and Issuer numbers
package index

import (
	"os"

	"github.com/park-jun-woo/silk/pkg/seg"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// ingest turns documents into segment rows. Tag i's rows go to buckets[i]
// and are written as a sorted run when sorted[i], else as an appended
// segment; Batch bounds the rows held in memory per tag.
type ingest struct {
	dir     string
	m       *Manifest
	sp      spec.Spec
	table   []int
	buckets []*seg.Bucket
	sorted  []bool
	batch   uint64
	out     [][]SegRef
	read    int64
	written int64
	files   map[string]*os.File
	key     []byte
	buf     []byte
	packIx  map[string]uint32
	issIx   map[string]uint32
}
