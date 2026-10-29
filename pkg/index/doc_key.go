//ff:type feature=index type=model
//ff:what docKey — a version document's identity: Issuer, document ID, version ID
package index

// docKey matches archive documents with document table entries.
type docKey struct {
	issuer  string
	id      uint64
	version uint32
}
