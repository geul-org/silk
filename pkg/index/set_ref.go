//ff:type feature=index type=model
//ff:what SetRef — the generation of an archive set a manifest covers
package index

type SetRef struct {
	Name       string `json:"name"`
	Generation uint64 `json:"generation"`
	Manifest   string `json:"manifest"`
}
