//ff:func feature=seg type=builder control=sequence
//ff:what Bucket.Reset — drop the gathered rows (after they were written)
package seg

// Reset empties b.
func (b *Bucket) Reset() {
	b.parts, b.Rows = map[string][]byte{}, 0
}
