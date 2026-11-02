//ff:func feature=seg type=decoder control=sequence
//ff:what parseFooter — the shape and counts from a footer
package seg

import (
	"encoding/binary"
	"fmt"
)

func parseFooter(b []byte) (Info, uint64, uint64, error) {
	if len(b) != footerSize || string(b[:8]) != magic || string(b[footerSize-8:]) != magic || binary.BigEndian.Uint16(b[8:]) != version {
		return Info{}, 0, 0, fmt.Errorf("%w: footer", ErrFormat)
	}
	u := binary.BigEndian.Uint16
	i := Info{PartWidth: int(u(b[10:])), RowWidth: int(u(b[12:])), DocOff: int(u(b[14:]))}
	copy(i.Hash[:], b[32:64])
	return i, binary.BigEndian.Uint64(b[16:]), binary.BigEndian.Uint64(b[24:]), i.check()
}
