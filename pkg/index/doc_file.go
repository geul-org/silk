//ff:type feature=index type=model
//ff:what DocFile — one document table file
package index

type DocFile struct {
	File  string `json:"file"`
	First uint32 `json:"first"`
	Count uint32 `json:"count"`
}
