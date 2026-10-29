//ff:func feature=index type=decoder control=iteration dimension=1
//ff:what parseDocs — 40-byte document entries with their (TAG, count) pairs; pair ranges must lie inside the pairs area
package index

import (
	"encoding/binary"
	"fmt"
)

// parseDocs decodes entries e against the pairs area p.
func parseDocs(e, p []byte) ([]DocEntry, error) {
	be := binary.BigEndian
	out := make([]DocEntry, 0, len(e)/docSize)
	for ; len(e) > 0; e = e[docSize:] {
		d := DocEntry{Pack: be.Uint32(e), Issuer: be.Uint32(e[4:]), Off: be.Uint64(e[8:]), Len: be.Uint32(e[16:]),
			Version: be.Uint32(e[20:]), ID: be.Uint64(e[24:])}
		start, n := uint64(be.Uint32(e[32:])), uint64(be.Uint16(e[36:]))
		if (start+n)*pairSize > uint64(len(p)) || be.Uint16(e[38:]) != 0 {
			return nil, fmt.Errorf("%w: document entry %d", ErrFormat, len(out))
		}
		d.Counts = parsePairs(p[start*pairSize : (start+n)*pairSize])
		out = append(out, d)
	}
	return out, nil
}
