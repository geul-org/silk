//ff:type feature=index type=model
//ff:what docDiff — an archive against the document table: unchanged versions kept, new or moved versions added, indexed versions gone
package index

// docDiff is what an update changes.
type docDiff struct {
	kept   []keptDoc
	added  []liveDoc
	killed []uint32
}
