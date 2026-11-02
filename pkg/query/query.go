//ff:type feature=query type=model
//ff:what Query — conditions that must hold in one document, and how to answer
package query

type Query struct {
	Conds []Cond
	Mode  string
	Scan  bool
	Index string
}
