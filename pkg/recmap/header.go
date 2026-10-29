//ff:type feature=recmap type=model
//ff:what Header — a map file's identity and counts: Issuer, generation, manifest and pack SHA-256, pack size, keys, records, documents, body hash
package recmap

// Header is the fixed head of a map file. Issuer is canonical hex; Manifest
// and Pack are SHA-256 hex (the generation and data pack the map was built
// from). Sum is the SHA-256 of the body (everything after the header).
type Header struct {
	Issuer     string
	Generation uint64
	Manifest   string
	Pack       string
	PackSize   uint64
	Keys       uint32
	Records    uint64
	Docs       uint64
	Sum        [32]byte
}
