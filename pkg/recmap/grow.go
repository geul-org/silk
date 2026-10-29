//ff:func feature=recmap type=helper control=sequence
//ff:what grow — a reusable buffer resliced to n bytes, reallocated only when too small
package recmap

// grow returns buf[:n], allocating when cap(buf) < n.
func grow(buf []byte, n int) []byte {
	if cap(buf) < n {
		return make([]byte, n)
	}
	return buf[:n]
}
