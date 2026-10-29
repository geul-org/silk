//ff:func feature=recmap type=helper control=sequence
//ff:what bodyEnd — the file size a header implies (header, key table, document table, off[]), saturating on overflow
package recmap

import "math/bits"

// bodyEnd computes headerSize + K×24 + D×24 + N×4; an overflowing count
// gives MaxUint64, which no real file size equals.
func bodyEnd(h Header) uint64 {
	hiD, d := bits.Mul64(h.Docs, docSize)
	hiN, n := bits.Mul64(h.Records, offSize)
	s, c1 := bits.Add64(uint64(headerSize)+uint64(h.Keys)*keySize, d, 0)
	s, c2 := bits.Add64(s, n, 0)
	if hiD|hiN|c1|c2 != 0 {
		return ^uint64(0)
	}
	return s
}
