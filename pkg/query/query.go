//ff:type feature=query type=model
//ff:what Query — conditions a document must all meet
package query

// Query is a conjunction at document level.
type Query struct {
	Conds []Cond
}
