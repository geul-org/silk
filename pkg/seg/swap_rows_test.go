//ff:func feature=seg type=helper control=sequence
//ff:what swapRows — test helper: a copy of b with the w-byte rows at offsets i and j swapped
package seg

// swapRows breaks row order on purpose.
func swapRows(b []byte, i, j, w int) []byte {
	out := append([]byte{}, b...)
	copy(out[i:i+w], b[j:j+w])
	copy(out[j:j+w], b[i:i+w])
	return out
}
