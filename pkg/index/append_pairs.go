//ff:func feature=index type=encoder control=iteration dimension=1
//ff:what appendPairs — (TAG WORD u16, count u32) pairs, big-endian
package index

import "encoding/binary"

// appendPairs encodes a document's record counts.
func appendPairs(b []byte, cs []TagCount) []byte {
	for _, c := range cs {
		b = binary.BigEndian.AppendUint32(binary.BigEndian.AppendUint16(b, c.Tag), c.N)
	}
	return b
}
