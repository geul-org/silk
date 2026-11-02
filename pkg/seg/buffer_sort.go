//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what Buffer.sort — order the sort index by partition and row
package seg

import (
	"bytes"
	"slices"
)

func (b *Buffer) sort() {
	n := b.Rows()
	if cap(b.idx) < n {
		b.idx = make([]uint32, n)
	}
	b.idx = b.idx[:n]
	for i := range b.idx {
		b.idx[i] = uint32(i)
	}
	slices.SortFunc(b.idx, func(x, y uint32) int { return bytes.Compare(b.at(x), b.at(y)) })
}
