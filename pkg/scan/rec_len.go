//ff:func feature=scan type=decoder control=selection
//ff:what RecLen — a record's byte length from the TAG WORD and Length extension at b; 0 when b is too short to tell, -1 when the prefix is not 000
//ff:why core.DecodeRecord needs the record as []Word; the hot loop frames raw bytes in place. A byte-level framing helper is a candidate for gogeul core.
package scan

import "encoding/binary"

// RecLen reads `000K KLLL LCCC CCCC` (record.md): LLLL 0-13 is 1-14 words,
// 14 a 1-word extension, 15 a 2-word extension, both counting every word.
func RecLen(b []byte) int {
	if len(b) < 2 {
		return 0
	}
	t := binary.BigEndian.Uint16(b)
	if t>>13 != 0 {
		return -1
	}
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
