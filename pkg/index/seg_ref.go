//ff:type feature=index type=model
//ff:what SegRef — a segment file of a section: relative path, rows (partitions and other), bytes
package index

// SegRef names a segment in a manifest.
type SegRef struct {
	File  string `json:"file"`
	Rows  uint64 `json:"rows"`
	Bytes int64  `json:"bytes"`
}
