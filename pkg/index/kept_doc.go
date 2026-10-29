//ff:type feature=index type=model
//ff:what keptDoc — an indexed version still live and unchanged: its sequence number and its place in the archive
package index

// keptDoc is a document an update does not re-index (except for rebuilt tags).
type keptDoc struct {
	seq uint32
	doc liveDoc
}
