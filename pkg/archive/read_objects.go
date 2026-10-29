//ff:func feature=archive type=decoder control=iteration dimension=1
//ff:what readObjects — OBJECTS body: u32 count, then 48-byte entries (kind, reserved, size, SHA-256)
package archive

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// readObjects decodes an OBJECTS section body; its length must be exact.
func readObjects(sec []byte) ([]Object, error) {
	if len(sec) < 4 {
		return nil, fmt.Errorf("%w: OBJECTS truncated", ErrManifest)
	}
	n := uint64(binary.LittleEndian.Uint32(sec))
	if uint64(len(sec)-4) != n*objectSize {
		return nil, fmt.Errorf("%w: OBJECTS length", ErrManifest)
	}
	out := make([]Object, 0, n)
	for e := sec[4:]; len(e) > 0; e = e[objectSize:] {
		out = append(out, Object{Kind: e[0], Size: binary.LittleEndian.Uint64(e[8:]), Hash: hex.EncodeToString(e[16:48])})
	}
	return out, nil
}
