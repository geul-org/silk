//ff:func feature=grammar type=decoder control=selection
//ff:what fieldLen — the byte length of field f at q; -1 when it does not fit
package grammar

func fieldLen(rec []byte, f Field, q, tail int) int {
	if q > len(rec) {
		return -1
	}
	n := 2 * f.Width
	switch f.Width {
	case 0:
		n = RecLen(rec[q:])
		if n == 0 {
			return -1
		}
	case RestWidth:
		n = len(rec) - tail - q
	}
	if n < 0 || n > len(rec)-q {
		return -1
	}
	return n
}
