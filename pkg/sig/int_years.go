//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what intYears — a minimal-width integer payload of up to 8 bytes
package sig

func intYears(p []byte, signed bool) (int64, int64, bool) {
	if len(p) == 0 || len(p) > 8 {
		return 0, 0, false
	}
	var n int64
	if signed && p[0]&0x80 != 0 {
		n = -1
	}
	for _, b := range p {
		n = n<<8 | int64(b)
	}
	return n, n, n >= 0 || signed
}
