//ff:type feature=archive type=model
//ff:what Object — one manifest OBJECTS entry: kind (1 data pack, 2 index object), size and SHA-256 hex
package archive

// Object is an entry of a manifest's OBJECTS section; its number is its
// position there.
type Object struct {
	Kind byte
	Size uint64
	Hash string
}
