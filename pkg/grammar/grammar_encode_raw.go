//ff:func feature=grammar type=encoder control=sequence
//ff:what Grammar.encodeRaw — a value of a raw field: "text" its bytes, 0x hex, or a decimal number in the field's width
package grammar

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func (g *Grammar) encodeRaw(l Leaf, tok string) ([][]byte, error) {
	if strings.HasPrefix(tok, `"`) {
		s, err := strconv.Unquote(tok)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrSyntax, tok)
		}
		return [][]byte{[]byte(s)}, nil
	}
	if strings.HasPrefix(tok, "0x") {
		b, err := hex.DecodeString(tok[2:])
		return [][]byte{b}, wrapHex(err)
	}
	n, err := strconv.ParseUint(tok, 10, 64)
	if err != nil || l.Field.Width == RestWidth {
		return nil, fmt.Errorf("%w: %s", ErrSyntax, tok)
	}
	b := binary.BigEndian.AppendUint64(nil, n)
	w := 2 * l.Field.Width
	return [][]byte{append(make([]byte, max(0, w-8)), b[max(0, 8-w):]...)}, nil
}
