//ff:type feature=index type=model
//ff:what PackRef — one data pack the document table numbers into
package index

type PackRef struct {
	Set  string `json:"set"`
	Hash string `json:"hash"`
}
