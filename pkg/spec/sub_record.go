//ff:func feature=spec type=decoder control=sequence
//ff:what subRecord — the language record at the start of a UTF8-LANGUAGE-TEXT Value's payload; false for any other record (a LOCAL Reference needs resolving)
package spec

import "encoding/binary"

// subRecord reads the first embedded record of f's payload when f is a
// UTF8-LANGUAGE-TEXT Value (0x4A or 0x4B).
func subRecord(f []byte) ([]byte, bool) {
	t := binary.BigEndian.Uint16(f)
	if uint8(t>>11)&3 != kindValue || uint8(t&0x7F)&^1 != codecLangTextEven {
		return nil, false
	}
	q := headLen(f)
	e, ok := nextField(f, q)
	if !ok {
		return nil, false
	}
	return f[q:e], true
}
