//ff:func feature=query type=encoder control=sequence
//ff:what encodeHex — hex digits as record bytes; they must decode as exactly one record
package query

import (
	"encoding/hex"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// encodeHex takes a record written out in hex (as geul show prints them).
func encodeHex(h string) ([]byte, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return nil, fmt.Errorf("%w: 0x%s: %v", ErrSyntax, h, err)
	}
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, fmt.Errorf("%w: 0x%s: %v", ErrSyntax, h, err)
	}
	if _, n, err := core.DecodeRecord(ws); err != nil || n != len(ws) {
		return nil, fmt.Errorf("%w: 0x%s is not one record", ErrSyntax, h)
	}
	return b, nil
}
