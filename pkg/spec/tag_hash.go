//ff:func feature=spec type=encoder control=sequence
//ff:what Tag.Hash — SHA-256 of one tag's canonical JSON; an index rebuilds only the sections whose tag hash changed
package spec

import (
	"crypto/sha256"
	"encoding/json"
)

// Hash identifies one tag's declaration.
func (t Tag) Hash() [32]byte {
	b, _ := json.Marshal(t)
	return sha256.Sum256(b)
}
