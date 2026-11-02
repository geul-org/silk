//ff:func feature=index type=reader control=sequence
//ff:what readBitmap — a live file back to a Bitmap; hash and size checked
package index

import (
	"fmt"
	"os"
)

func readBitmap(path string) (Bitmap, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Bitmap{}, err
	}
	_, n, body, err := readTail(path, b, liveMagic)
	if err != nil {
		return Bitmap{}, err
	}
	if uint64(len(body)) != (uint64(n)+7)/8 {
		return Bitmap{}, fmt.Errorf("%w: %s size", ErrFormat, path)
	}
	return Bitmap{B: body, N: n}, nil
}
