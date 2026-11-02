//ff:func feature=index type=encoder control=sequence
//ff:what appendTail — the 56-byte tail of doc and live files: magic, first, count, reserved, body SHA-256
package index

import (
	"crypto/sha256"
	"encoding/binary"
)

func appendTail(out []byte, magic string, first, count uint32, body []byte) []byte {
	out = append(out, magic...)
	out = binary.BigEndian.AppendUint32(out, first)
	out = binary.BigEndian.AppendUint32(out, count)
	out = append(out, make([]byte, 8)...)
	s := sha256.Sum256(body)
	return append(out, s[:]...)
}
