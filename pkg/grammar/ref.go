//ff:func feature=grammar type=encoder control=selection
//ff:what Ref — the record of a Qn, Pn or #n token
package grammar

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

func Ref(tok string) ([]byte, error) {
	if len(tok) < 2 {
		return nil, fmt.Errorf("%w: %q", ErrSyntax, tok)
	}
	var c uint8
	switch tok[0] {
	case 'Q':
		c = core.RefWikidataQID
	case 'P':
		c = core.RefWikidataPID
	case '#':
		c = core.RefLocal
	default:
		return nil, fmt.Errorf("%w: %q is not Qn, Pn or #n", ErrSyntax, tok)
	}
	n, err := strconv.ParseUint(tok[1:], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %q needs a decimal number", ErrSyntax, tok)
	}
	return RecordOf(core.Record{Kind: core.KindReference, Codec: c << 3, Payload: core.EncodeUintValue(n).Payload}), nil
}
