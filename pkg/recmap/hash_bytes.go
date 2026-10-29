//ff:func feature=recmap type=encoder control=sequence
//ff:what hashBytes — a SHA-256 as 64 lowercase hex digits (pack object naming) to its 32 bytes
package recmap

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// hashBytes reads a pack object name hash.
func hashBytes(s string) ([32]byte, error) {
	var out [32]byte
	b, err := hex.DecodeString(s)
	if err != nil || len(b) != 32 || strings.ToLower(s) != s {
		return out, fmt.Errorf("%w: SHA-256 %q", ErrFormat, s)
	}
	copy(out[:], b)
	return out, nil
}
