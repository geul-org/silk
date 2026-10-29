//ff:func feature=spec type=decoder control=selection
//ff:what headLen — bytes of a record's TAG WORD and Length extension (2, 4 or 6) from its Length code; 0 when rec has no TAG WORD
package spec

import "encoding/binary"

// headLen tells where a record's payload starts.
func headLen(rec []byte) int {
	if len(rec) < 2 {
		return 0
	}
	switch (binary.BigEndian.Uint16(rec) >> 7) & 0xF {
	case 0xE:
		return 4
	case 0xF:
		return 6
	default:
		return 2
	}
}
