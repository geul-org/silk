//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what get — read w bits at bit offset off
package sig

func get(b []byte, off, w int) uint64 {
	var v uint64
	for i := 0; i < w; i++ {
		p := off + i
		v = v<<1 | uint64(b[p/8]>>(7-p%8)&1)
	}
	return v
}
