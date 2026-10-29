//ff:type feature=index type=model
//ff:what Index — the opened maps of an index directory, each with its data pack
package index

// Index is what Open returns; Close releases its files.
type Index struct {
	Dir   string
	Parts []Part
}
