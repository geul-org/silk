//ff:func feature=seg type=encoder control=iteration dimension=1
//ff:what appendTable — partition table entries (key slot, start u64, count u64, big-endian) appended to b
package seg

import "encoding/binary"

// appendTable returns a new slice (b is not modified in place).
func appendTable(b []byte, parts []Part) []byte {
	out := append([]byte{}, b...)
	for _, p := range parts {
		out = binary.BigEndian.AppendUint64(binary.BigEndian.AppendUint64(append(out, p.Key...), p.Start), p.Count)
	}
	return out
}
