//ff:type feature=index type=model
//ff:what job — one document to read
package index

type job struct {
	seq   uint32
	entry DocEntry
	isNew bool
}
