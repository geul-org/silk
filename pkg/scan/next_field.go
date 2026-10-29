//ff:func feature=scan type=decoder control=sequence
//ff:what nextField — the end of the embedded record starting at p, false when it does not fit inside rec
package scan

// nextField frames the field record at rec[p:].
func nextField(rec []byte, p int) (int, bool) {
	if p >= len(rec) {
		return p, false
	}
	n := RecLen(rec[p:])
	if n <= 0 || n > len(rec)-p {
		return p, false
	}
	return p + n, true
}
