//ff:func feature=seg type=decoder control=sequence
//ff:what parseFooter — the 128-byte footer back to an Info; wrong magic, version, flags, widths or nonzero padding is ErrFormat
package seg

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// parseFooter reads b (exactly footerSize bytes).
func parseFooter(b []byte) (Info, error) {
	be := binary.BigEndian
	if len(b) != footerSize || string(b[:8]) != magic || string(b[120:]) != magic || be.Uint16(b[8:]) != version {
		return Info{}, fmt.Errorf("%w: footer", ErrFormat)
	}
	flags := be.Uint16(b[10:])
	i := Info{Tag: be.Uint16(b[12:]), PartWidth: int(be.Uint16(b[14:])), SortWidth: int(be.Uint16(b[16:])),
		Sorted: flags&flagSorted != 0, Parts: be.Uint32(b[20:]), Rows: be.Uint64(b[24:]), Other: be.Uint64(b[32:])}
	copy(i.Spec[:], b[40:72])
	copy(i.Sum[:], b[72:104])
	if flags&^flagSorted != 0 || be.Uint16(b[18:]) != 0 || i.PartWidth > 64 || i.SortWidth > 64 || !bytes.Equal(b[104:120], make([]byte, 16)) {
		return Info{}, fmt.Errorf("%w: footer fields", ErrFormat)
	}
	return i, nil
}
