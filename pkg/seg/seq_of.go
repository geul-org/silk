//ff:func feature=seg type=decoder control=sequence
//ff:what SeqOf — a row's document sequence number and record offset (its last 8 bytes, big-endian)
package seg

import "encoding/binary"

// SeqOf reads the tail of a row.
func SeqOf(row []byte) (uint32, uint32) {
	n := len(row)
	return binary.BigEndian.Uint32(row[n-8:]), binary.BigEndian.Uint32(row[n-4:])
}
