//ff:func feature=index type=helper control=iteration dimension=1
//ff:what allLive — a bitmap with sequence numbers 0..n-1 live
package index

// allLive is the bitmap of a fresh build.
func allLive(n uint32) Bitmap {
	var b Bitmap
	for i := range n {
		b.Put(i, true)
	}
	return b
}
