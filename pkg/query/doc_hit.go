//ff:type feature=query type=model
//ff:what docHit — a document meeting the conditions so far, with the Edge Ids of its matching records
package query

// docHit is one row of a merge join.
type docHit struct {
	doc   uint32
	edges []uint64
}
