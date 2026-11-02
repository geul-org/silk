//ff:type feature=archive type=model
//ff:what Object — one object of a pack manifest: kind, size and SHA-256 hex
package archive

type Object struct {
	Kind byte
	Size uint64
	Hash string
}
