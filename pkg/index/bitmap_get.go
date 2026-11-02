//ff:func feature=index type=helper control=sequence
//ff:what Bitmap.Get — whether a document number is live (false beyond N)
package index

func (b Bitmap) Get(i uint32) bool {
	return i < b.N && b.B[i/8]&(1<<(i%8)) != 0
}
