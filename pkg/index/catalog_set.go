//ff:type feature=index type=model
//ff:what CatalogSet — one set as the catalog records it: folder name, generation, manifest SHA-256 hex
package index

// CatalogSet identifies a set generation.
type CatalogSet struct {
	Name       string `json:"name"`
	Generation uint64 `json:"generation"`
	Manifest   string `json:"manifest"`
}
