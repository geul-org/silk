//ff:func feature=query type=encoder control=selection
//ff:what encodeToken — Qn, Pn or #n as the canonical bytes of a WIKIDATA-QID, WIKIDATA-PID or LOCAL Reference record; * is nil (any)
//ff:why gogeul notation reads these tokens but keeps parseRecord unexported; an exported token reader is a candidate for gogeul.
package query

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// encodeToken builds the record as gogeul's notation does (Reference codec
// CCCC<<3, minimal-word target ID) and serializes it.
func encodeToken(tok string) ([]byte, error) {
	var c uint8
	switch {
	case tok == "*":
		return nil, nil
	case tok == "":
		return nil, fmt.Errorf("%w: empty token", ErrSyntax)
	case tok[0] == 'Q':
		c = core.RefWikidataQID
	case tok[0] == 'P':
		c = core.RefWikidataPID
	case tok[0] == '#':
		c = core.RefLocal
	default:
		return nil, fmt.Errorf("%w: %q is not Qn, Pn, #n or *", ErrSyntax, tok)
	}
	n, err := strconv.ParseUint(tok[1:], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %q needs a decimal number", ErrSyntax, tok)
	}
	ws, _ := core.EncodeRecord(core.Record{Kind: core.KindReference, Codec: c << 3, Payload: core.EncodeUintValue(n).Payload})
	return core.PutWords(ws), nil
}
