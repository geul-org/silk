//ff:type feature=index type=model
//ff:what PackRef — a data pack the document table points into: its set and SHA-256 hex
package index

// PackRef is `<set>/o/<Hash>.gpk`.
type PackRef struct {
	Set  string `json:"set"`
	Hash string `json:"hash"`
}
