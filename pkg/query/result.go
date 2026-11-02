//ff:type feature=query type=model
//ff:what Result — one query's hits and the plans that made them
package query

type Result struct {
	Hits    []Hit
	Plans   []string
	Decoded int
}
