//ff:func feature=grammar type=encoder control=selection
//ff:what Grammar.Encode — the byte strings a value token stands for at a field: more than one when the grammar calls them equal (SHORT-CODE and BCP 47)
package grammar

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func (g *Grammar) Encode(l Leaf, tok string) ([][]byte, error) {
	if l.Field.Width != 0 {
		return g.encodeRaw(l, tok)
	}
	switch {
	case tok == "":
		return nil, fmt.Errorf("%w: empty value", ErrSyntax)
	case tok[0] == '"':
		return g.encodeText(tok)
	case tok[0] == '@':
		rs, err := g.langRecords(tok[1:], false)
		return recordsOf(rs), err
	case strings.HasPrefix(tok, "0x"):
		b, err := hex.DecodeString(tok[2:])
		return [][]byte{b}, wrapHex(err)
	}
	b, err := Ref(tok)
	return [][]byte{b}, err
}
