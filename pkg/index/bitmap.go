//ff:type feature=index type=model
//ff:what Bitmap — live document versions by sequence number, one bit each
package index

// Bitmap bit i (byte i/8, bit i%8) is set when sequence i is live. N is the
// number of sequence numbers it covers.
type Bitmap struct {
	B []byte
	N uint32
}
