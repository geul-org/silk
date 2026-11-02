//ff:type feature=index type=model
//ff:what docMap — one mapped table file
package index

type docMap struct {
	first, count uint32
	raw          []byte
}
