//ff:func feature=archive type=decoder control=iteration dimension=1
//ff:what sectionOf — the body of one manifest section by kind, bounds checked
package archive

import (
	"encoding/binary"
	"fmt"
)

// sectionOf scans the n-entry section table after the manifest header for kind.
func sectionOf(b []byte, n int, kind uint16) ([]byte, error) {
	if n < 0 || len(b) < manifestHead+n*sectionSize {
		return nil, fmt.Errorf("%w: section table truncated", ErrManifest)
	}
	for i := 0; i < n; i++ {
		e := b[manifestHead+i*sectionSize:]
		if binary.LittleEndian.Uint16(e) != kind {
			continue
		}
		off, size := binary.LittleEndian.Uint64(e[8:]), binary.LittleEndian.Uint64(e[16:])
		if off > uint64(len(b)) || size > uint64(len(b))-off {
			return nil, fmt.Errorf("%w: section 0x%04x out of bounds", ErrManifest, kind)
		}
		return b[off : off+size], nil
	}
	return nil, fmt.Errorf("%w: no section 0x%04x", ErrManifest, kind)
}
