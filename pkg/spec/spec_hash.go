//ff:func feature=spec type=encoder control=sequence
//ff:what Spec.Hash — SHA-256 hex of the canonical spec bytes
package spec

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash identifies the spec content.
func (s Spec) Hash() string {
	h := sha256.Sum256(s.Encode())
	return hex.EncodeToString(h[:])
}
