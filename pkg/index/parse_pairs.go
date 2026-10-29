//ff:func feature=index type=decoder control=iteration dimension=1
//ff:what parsePairs — (TAG WORD, count) pairs back to TagCounts
package index

import "encoding/binary"

// parsePairs decodes 6-byte pairs.
func parsePairs(b []byte) []TagCount {
	out := make([]TagCount, 0, len(b)/pairSize)
	for ; len(b) >= pairSize; b = b[pairSize:] {
		out = append(out, TagCount{Tag: binary.BigEndian.Uint16(b), N: binary.BigEndian.Uint32(b[2:])})
	}
	return out
}
