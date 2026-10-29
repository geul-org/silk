//ff:type feature=query type=model
//ff:what Triple — one TRIPLE condition as written and as the canonical record bytes of its Predicate and Object (nil = any)
package query

// Triple is a condition `P=O`; Pred and Obj are encoded GEUL records.
type Triple struct {
	Text string
	Pred []byte
	Obj  []byte
}
