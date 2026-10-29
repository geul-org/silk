//ff:type feature=archive type=model
//ff:what Doc — one version document: Issuer hex, document ID, version ID and its raw bytes' place in a data pack
package archive

// Doc is the version document `<Issuer>/<ID>/<Version>.geul`, stored raw at
// [Off, Off+Len) of its data pack.
type Doc struct {
	Issuer  string
	ID      uint64
	Version uint32
	Off     uint64
	Len     uint64
}
