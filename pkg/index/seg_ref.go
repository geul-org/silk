//ff:type feature=index type=model
//ff:what SegRef — one segment of a section
package index

type SegRef struct {
	File  string `json:"file"`
	Rows  uint64 `json:"rows"`
	Bytes uint64 `json:"bytes"`
}
