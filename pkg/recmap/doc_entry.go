//ff:type feature=recmap type=model
//ff:what DocEntry — one version document of the map: its byte range in the data pack, document ID and version ID
package recmap

// DocEntry is a document at [Off, Off+Len) of the data pack.
type DocEntry struct {
	Off     uint32
	Len     uint32
	ID      uint64
	Version uint32
}
