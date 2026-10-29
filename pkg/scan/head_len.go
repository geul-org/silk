//ff:func feature=scan type=decoder control=selection
//ff:what headLen — bytes of a record's TAG WORD and Length extension (2, 4 or 6) from its Length code
package scan

import "encoding/binary"

// headLen assumes rec holds at least the TAG WORD (as a framed record does).
func headLen(rec []byte) int {
	switch (binary.BigEndian.Uint16(rec) >> 7) & 0xF {
	case 0xE:
		return 4
	case 0xF:
		return 6
	default:
		return 2
	}
}
