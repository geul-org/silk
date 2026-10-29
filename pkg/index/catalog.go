//ff:type feature=index type=model
//ff:what Catalog — index.json: format version, the set generations the maps were built from, and the map files (relative, slash-separated)
package index

// Catalog is what Build writes last and Open reads first.
type Catalog struct {
	Version int          `json:"version"`
	Sets    []CatalogSet `json:"sets"`
	Maps    []string     `json:"maps"`
}
