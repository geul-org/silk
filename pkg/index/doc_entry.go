//ff:type feature=index type=model
//ff:what DocEntry — one document number: pack, Issuer, stored location and identity; Enc 0xFF marks a deletion step
package index

type DocEntry struct {
	Pack    uint32
	Issuer  uint32
	Off     uint64
	Len     uint64
	ID      uint64
	Version uint32
	Enc     byte
}
