//ff:func feature=recmap type=decoder control=iteration dimension=1
//ff:what parseDocs — the document table: offsets ascending without overlap, reserved bytes zero
package recmap

import (
	"encoding/binary"
	"fmt"
)

// parseDocs decodes 24-byte entries.
func parseDocs(b []byte) ([]DocEntry, error) {
	le := binary.LittleEndian
	out := make([]DocEntry, 0, len(b)/docSize)
	end := uint64(0)
	for e := b; len(e) > 0; e = e[docSize:] {
		d := DocEntry{Off: le.Uint32(e), Len: le.Uint32(e[4:]), ID: le.Uint64(e[8:]), Version: le.Uint32(e[16:])}
		if uint64(d.Off) < end || le.Uint32(e[20:]) != 0 {
			return nil, fmt.Errorf("%w: document table entry %d", ErrFormat, len(out))
		}
		out = append(out, d)
		end = uint64(d.Off) + uint64(d.Len)
	}
	return out, nil
}
