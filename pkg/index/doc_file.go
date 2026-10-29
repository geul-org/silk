//ff:type feature=index type=model
//ff:what DocFile — a document table file: relative path, first document sequence number, count
package index

// DocFile holds sequence numbers [First, First+Count).
type DocFile struct {
	File  string `json:"file"`
	First uint32 `json:"first"`
	Count uint32 `json:"count"`
}
