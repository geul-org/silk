//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what Buffer.Unique — keep one of each byte-equal row
package seg

import "bytes"

func (b *Buffer) Unique() {
	b.sort()
	out := make([]byte, 0, len(b.data))
	w := b.pw + b.rw
	for _, i := range b.idx {
		r := b.at(i)
		if len(out) == 0 || !bytes.Equal(out[len(out)-w:], r) {
			out = append(out, r...)
		}
	}
	b.data = append(b.data[:0], out...)
}
