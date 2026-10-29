//ff:func feature=index type=helper control=sequence
//ff:what Bitmap.Get — whether a sequence number is live (false beyond N)
package index

// Get reads bit i.
func (b Bitmap) Get(i uint32) bool {
	return i < b.N && b.B[i/8]&(1<<(i%8)) != 0
}
