//ff:func feature=index type=decoder control=sequence
//ff:what readTail — first, count and body of a tailed file, hash checked
package index

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func readTail(path string, b []byte, magic string) (uint32, uint32, []byte, error) {
	if len(b) < tailSize || string(b[len(b)-tailSize:len(b)-tailSize+8]) != magic {
		return 0, 0, nil, fmt.Errorf("%w: %s", ErrFormat, path)
	}
	t := b[len(b)-tailSize:]
	body := b[:len(b)-tailSize]
	if s := sha256.Sum256(body); !bytes.Equal(s[:], t[24:]) {
		return 0, 0, nil, fmt.Errorf("%w: %s hash", ErrFormat, path)
	}
	return binary.BigEndian.Uint32(t[8:]), binary.BigEndian.Uint32(t[12:]), body, nil
}
