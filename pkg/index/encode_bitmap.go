//ff:func feature=index type=encoder control=sequence
//ff:what encodeBitmap — the live file: magic, version, bit count, body SHA-256, then the bits
package index

import (
	"crypto/sha256"
	"encoding/binary"
)

// encodeBitmap lays out docs/index.md "Live bitmap".
func encodeBitmap(b Bitmap) []byte {
	out := append(make([]byte, 0, liveHead+len(b.B)), liveMagic...)
	out = binary.BigEndian.AppendUint16(out, fileVersion)
	out = append(out, 0, 0)
	out = binary.BigEndian.AppendUint32(out, b.N)
	out = append(out, make([]byte, 8)...)
	s := sha256.Sum256(b.B)
	return append(append(out, s[:]...), b.B...)
}
