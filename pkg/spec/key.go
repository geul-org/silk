//ff:type feature=spec type=model
//ff:what Key — a key's field path ([field] or [field, 0] for a language record) and its slot width in bytes; no path is no key
package spec

// Key names a record field by position and how many of its leading bytes the
// index keeps (zero padded). An empty Path means the TAG has no such key.
type Key struct {
	Path  []int `json:"path,omitempty"`
	Width int   `json:"width,omitempty"`
}
