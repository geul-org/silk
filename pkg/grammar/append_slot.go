//ff:func feature=grammar type=encoder control=sequence
//ff:what AppendSlot — a key slot of width w: a record's or raw field's first w bytes, a rest field's first w-1 and a last byte 00 whole or FF cut; zero when missing
package grammar

func AppendSlot(dst []byte, v Val, ok bool, w int) []byte {
	start := len(dst)
	dst = append(dst, make([]byte, w)...)
	if !ok {
		return dst
	}
	n := copy(dst[start:start+w], v.B)
	if v.Rest {
		n = copy(dst[start:start+w-1], v.B)
		dst[start+w-1] = 0
		clear(dst[start+n : start+w])
		if len(v.B) > w-1 {
			dst[start+w-1] = 0xFF
		}
	}
	return dst
}
