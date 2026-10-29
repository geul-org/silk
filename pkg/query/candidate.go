//ff:type feature=query type=model
//ff:what candidate — a record a segment lookup returned: document sequence number and offset in the document
package query

// candidate is checked exactly before it counts.
type candidate struct {
	doc uint32
	off uint32
}
