//ff:func feature=seg type=builder control=sequence
//ff:what Bucket.Add — one row under its partition key slot
package seg

// Add appends row (sort slot, document, offset) to partition key.
func (b *Bucket) Add(key, row []byte) {
	b.parts[string(key)] = append(b.parts[string(key)], row...)
	b.Rows++
}
