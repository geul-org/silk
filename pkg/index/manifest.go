//ff:type feature=index type=model
//ff:what Manifest — one index generation: spec hash, archive sets, packs, Issuers, document files, live bitmap, sections, next file number
package index

// Manifest is `manifest/<generation>.json`; CURRENT names the live one.
type Manifest struct {
	Version    int       `json:"version"`
	Generation uint64    `json:"generation"`
	Spec       string    `json:"spec"`
	Sets       []SetRef  `json:"sets"`
	Packs      []PackRef `json:"packs"`
	Issuers    []string  `json:"issuers"`
	Docs       []DocFile `json:"docs"`
	Live       string    `json:"live"`
	Sections   []Section `json:"sections"`
	NextFile   uint64    `json:"next_file"`
}
