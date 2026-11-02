//ff:func feature=index type=helper control=iteration dimension=1
//ff:what fingerprint.add — count one item
package index

import (
	"crypto/sha256"
	"encoding/binary"
)

func (p *fingerprint) add(b ...[]byte) {
	h := sha256.New()
	for _, x := range b {
		h.Write(x)
	}
	p.n++
	p.sum += binary.BigEndian.Uint64(h.Sum(nil))
}
