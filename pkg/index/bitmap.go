//ff:type feature=index type=model
//ff:what Bitmap — live document numbers, one bit each
package index

type Bitmap struct {
	B []byte
	N uint32
}
