//ff:type feature=index type=model
//ff:what SetRef — an archive set generation the index covers: folder name, generation, manifest SHA-256 hex
package index

// SetRef identifies a set generation.
type SetRef struct {
	Name       string `json:"name"`
	Generation uint64 `json:"generation"`
	Manifest   string `json:"manifest"`
}
