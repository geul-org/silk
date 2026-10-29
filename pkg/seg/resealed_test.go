//ff:func feature=seg type=helper control=sequence
//ff:what resealed — test helper: f over edited bytes with the footer's body hash recomputed, so only the structure checks can fail
package seg

import (
	"bytes"
	"crypto/sha256"
)

// resealed returns a File reading b (same size as f) with a matching Sum.
func resealed(f *File, b []byte) *File {
	g := *f
	g.Info.Sum = sha256.Sum256(b[:len(b)-footerSize])
	g.r = bytes.NewReader(b)
	return &g
}
