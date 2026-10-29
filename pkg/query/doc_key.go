//ff:type feature=query type=model
//ff:what docKey — a version document's identity while collecting hits
package query

// docKey identifies a version document.
type docKey struct {
	issuer  string
	doc     uint64
	version uint32
}
