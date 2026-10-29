//ff:func feature=recmap type=reader control=iteration dimension=1
//ff:what File.offs — count entries of off[] from start, read with one ReadAt and decoded little-endian
package recmap

import (
	"encoding/binary"
	"fmt"
)

// offs reads off[start : start+count].
func (f *File) offs(start, count uint64) ([]uint32, error) {
	at := headerSize + uint64(f.Header.Keys)*keySize + f.Header.Docs*docSize + start*offSize
	b := make([]byte, count*offSize)
	if _, err := f.r.ReadAt(b, int64(at)); err != nil {
		return nil, fmt.Errorf("%w: off[]: %v", ErrFormat, err)
	}
	out := make([]uint32, count)
	for i := range out {
		out[i] = binary.LittleEndian.Uint32(b[i*offSize:])
	}
	return out, nil
}
