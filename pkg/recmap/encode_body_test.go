//ff:func feature=recmap type=encoder control=sequence
//ff:what TestEncodeBody — table sizes and little-endian fields at their documented positions
package recmap

import (
	"encoding/binary"
	"testing"
)

func TestEncodeBody(t *testing.T) {
	m := &Map{Keys: []KeyRange{{Key: 0x100a, Start: 0, Count: 2}}, Docs: []DocEntry{{Off: 64, Len: 8, ID: 7, Version: 1}}, Offs: []uint32{64, 66}}
	b := encodeBody(m)
	le := binary.LittleEndian
	if len(b) != keySize+docSize+2*offSize || le.Uint16(b) != 0x100a || le.Uint64(b[16:]) != 2 ||
		le.Uint32(b[24:]) != 64 || le.Uint64(b[32:]) != 7 || le.Uint32(b[40:]) != 1 || le.Uint32(b[52:]) != 66 {
		t.Fatal(b)
	}
}
