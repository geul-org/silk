//ff:type feature=index type=model
//ff:what DocEntry — one indexed document version: pack and Issuer numbers, byte range in the pack, document and version IDs, records per TAG kind
package index

// DocEntry is the document table row of one sequence number. Counts gives
// the number of top-level records of each Kind·Codec, so removing the
// version tells each section how many of its rows died.
type DocEntry struct {
	Pack    uint32
	Issuer  uint32
	Off     uint64
	Len     uint32
	ID      uint64
	Version uint32
	Counts  []TagCount
}
