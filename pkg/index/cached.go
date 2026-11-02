//ff:type feature=index type=model
//ff:what cached — one cached document
package index

type cached struct {
	seq uint32
	b   []byte
}
