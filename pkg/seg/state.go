//ff:type feature=seg type=model
//ff:what State — a paused Writer at a partition boundary: target path, bytes written, hash state, table, other rows, last key
package seg

// State is what Checkpoint saves and ResumeWriter restores. It is plain data
// (JSON-encodable) so an interrupted merge can continue in another process.
type State struct {
	Path    string `json:"path"`
	Min     int    `json:"min"`
	Info    Info   `json:"info"`
	Parts   []Part `json:"parts"`
	Other   []byte `json:"other"`
	Key     []byte `json:"key"`
	Hash    []byte `json:"hash"`
	Written int64  `json:"written"`
}
