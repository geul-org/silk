//ff:type feature=index type=model
//ff:what Index — an opened, pinned index generation: manifest, spec, document table, live bitmap, segment files per section, data packs (opened on demand)
package index

import (
	"os"
	"sync"

	"github.com/geul-org/silk/pkg/seg"
	"github.com/geul-org/silk/pkg/spec"
)

// Index is what Open returns. Its generation stays readable until Close,
// even when a writer commits newer ones.
type Index struct {
	Dir   string
	M     Manifest
	Spec  spec.Spec
	Docs  []DocEntry
	Live  Bitmap
	Segs  [][]*seg.File
	paths []string
	mu    sync.Mutex
	packs []*os.File
	pin   string
}
