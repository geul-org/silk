//ff:func feature=grammar type=decoder control=selection
//ff:what RecLen — a record's byte length from its TAG WORD and Length extension; 0 when b is too short or the prefix is not 000
package grammar

import "encoding/binary"

func RecLen(b []byte) int {
	if len(b) < 2 || b[0]>>5 != 0 {
		return 0
	}
	t := binary.BigEndian.Uint16(b)
	switch code := int(t>>7) & 0xF; {
	case code < 0xE:
		return 2 * (code + 1)
	case code == 0xE && len(b) >= 4:
		return 2 * int(binary.BigEndian.Uint16(b[2:]))
	case code == 0xF && len(b) >= 6:
		return 2 * int(binary.BigEndian.Uint32(b[2:]))
	default:
		return 0
	}
}
