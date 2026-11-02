//ff:type feature=index type=model
//ff:what Report — what a run did
package index

type Report struct {
	NewDocs  uint32
	ReadDocs uint64
	Rows     uint64
	Merges   int
}
