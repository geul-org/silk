//ff:func feature=grammar type=decoder control=sequence
//ff:what parseDocID — the 8 bytes of a codec spec document ID (`0x000000000000<TAG>`, LLLL = 0, not 0); false for other names
package grammar

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
)

func parseDocID(name string) ([]byte, bool) {
	if len(name) != 18 || !strings.HasPrefix(name, "0x") {
		return nil, false
	}
	b, err := hex.DecodeString(name[2:])
	if err != nil || binary.BigEndian.Uint64(b)>>16 != 0 {
		return nil, false
	}
	t := binary.BigEndian.Uint16(b[6:])
	return b, t != 0 && t&0x0780 == 0
}
