//ff:func feature=sig type=encoder control=iteration dimension=1
//ff:what encodeOneHot — set the bit of every listed value present; unlisted values set nothing
package sig

// encodeOneHot writes a OneHot field.
func encodeOneHot(s []uint64, f Field, vs []uint64) {
	for _, v := range vs {
		if b, ok := f.Bits[v]; ok {
			putBits(s, f.Bit+b, 1, 1)
		}
	}
}
