//ff:func feature=recmap type=rule control=iteration dimension=1
//ff:what checkDocs — documents ascend without overlap, lie inside the pack, and the pack fits u32 offsets
package recmap

import (
	"fmt"
	"math"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// checkDocs enforces what off[] u32 and the document table need.
func checkDocs(docs []archive.Doc, size uint64) error {
	if size > math.MaxUint32 {
		return fmt.Errorf("%w: pack of %d bytes exceeds u32 offsets", ErrFormat, size)
	}
	end := uint64(0)
	for _, d := range docs {
		if d.Off < end || d.Len > size || d.Off > size-d.Len {
			return fmt.Errorf("%w: document 0x%016x at %d+%d out of order or outside the pack", ErrFormat, d.ID, d.Off, d.Len)
		}
		end = d.Off + d.Len
	}
	return nil
}
