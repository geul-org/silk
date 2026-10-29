//ff:func feature=index type=reader control=sequence
//ff:what readBitmap — a live file back to a Bitmap; header, size and body SHA-256 checked
package index

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
)

// readBitmap reads path.
func readBitmap(path string) (Bitmap, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Bitmap{}, err
	}
	if len(b) < liveHead || string(b[:8]) != liveMagic || binary.BigEndian.Uint16(b[8:]) != fileVersion {
		return Bitmap{}, fmt.Errorf("%w: %s", ErrFormat, path)
	}
	n := binary.BigEndian.Uint32(b[12:])
	body := b[liveHead:]
	s := sha256.Sum256(body)
	if uint64(len(body)) != (uint64(n)+7)/8 || !bytes.Equal(s[:], b[24:56]) || !bytes.Equal(b[10:12], []byte{0, 0}) {
		return Bitmap{}, fmt.Errorf("%w: %s body", ErrFormat, path)
	}
	return Bitmap{B: body, N: n}, nil
}
