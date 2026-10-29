//ff:func feature=spec type=helper control=sequence
//ff:what hexBytes — test helper: space-separated hex words as bytes
package spec

import (
	"encoding/hex"
	"strings"
)

// hexBytes panics on bad input (test data is literal).
func hexBytes(s string) []byte {
	b, err := hex.DecodeString(strings.ReplaceAll(s, " ", ""))
	if err != nil {
		panic(err)
	}
	return b
}
