//ff:type feature=spec type=model
//ff:what Spec — the index specification: format version and the indexed TAG kinds
package spec

// Spec is what an index is built from; its hash goes in the index manifest.
type Spec struct {
	Version int   `json:"version"`
	Tags    []Tag `json:"tags"`
}
