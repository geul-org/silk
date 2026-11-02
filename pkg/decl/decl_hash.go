//ff:func feature=decl type=encoder control=sequence
//ff:what Decl.Hash — SHA-256 of the canonical JSON
package decl

import "crypto/sha256"

func (d Decl) Hash() [32]byte {
	return sha256.Sum256(d.Encode())
}
