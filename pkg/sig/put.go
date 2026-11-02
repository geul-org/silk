//ff:func feature=sig type=encoder control=iteration dimension=1
//ff:what put — write w bits of v at bit offset off, top bit first
package sig

func put(b []byte, off, w int, v uint64) {
	for i := 0; i < w; i++ {
		p, m := off+i, byte(0x80)>>((off+i)%8)
		b[p/8] &^= m
		if v>>(w-1-i)&1 == 1 {
			b[p/8] |= m
		}
	}
}
