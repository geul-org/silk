//ff:func feature=recmap type=encoder control=sequence
//ff:what issuerBytes — canonical Issuer hex (`0x`, 1/2/4/16 words) as its big-endian bytes
package recmap

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// issuerBytes reads `0x` and 4, 8, 16 or 64 lowercase hex digits.
func issuerBytes(s string) ([]byte, error) {
	h, ok := strings.CutPrefix(s, "0x")
	if n := len(h); !ok || (n != 4 && n != 8 && n != 16 && n != 64) || strings.ToLower(h) != h {
		return nil, fmt.Errorf("%w: Issuer %q", ErrFormat, s)
	}
	b, err := hex.DecodeString(h)
	if err != nil {
		return nil, fmt.Errorf("%w: Issuer %q", ErrFormat, s)
	}
	return b, nil
}
