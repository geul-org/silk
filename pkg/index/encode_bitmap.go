//ff:func feature=index type=encoder control=sequence
//ff:what encodeBitmap — the live file: bits, then magic, bit count, body SHA-256
package index

func encodeBitmap(b Bitmap) []byte {
	out := append(make([]byte, 0, len(b.B)+tailSize), b.B...)
	return appendTail(out, liveMagic, 0, b.N, b.B)
}
