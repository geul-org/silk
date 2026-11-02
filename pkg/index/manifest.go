//ff:type feature=index type=model
//ff:what Manifest — one generation of the index
package index

type Manifest struct {
	Version    int       `json:"version"`
	Generation uint64    `json:"generation"`
	Grammar    string    `json:"grammar,omitempty"`
	Sets       []SetRef  `json:"sets,omitempty"`
	Packs      []PackRef `json:"packs,omitempty"`
	Issuers    []string  `json:"issuers,omitempty"`
	Docs       []DocFile `json:"docs,omitempty"`
	Live       string    `json:"live,omitempty"`
	Paths      Section   `json:"paths"`
	Sections   []Section `json:"sections,omitempty"`
	MaxDoc     uint64    `json:"max_doc,omitempty"`
	NextFile   uint64    `json:"next_file"`
}
