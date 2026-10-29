//ff:type feature=index type=model
//ff:what TagCount — records of one TAG kind (Kind·Codec bits) in a document
package index

// TagCount pairs a TAG WORD with a record count.
type TagCount struct {
	Tag uint16
	N   uint32
}
