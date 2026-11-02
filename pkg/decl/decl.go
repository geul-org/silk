//ff:type feature=decl type=model
//ff:what Decl — one declared index: name, kind, TAG words, partition and key paths with widths, and for SILK kinds the definition text
package decl

type Decl struct {
	Name      string   `json:"name"`
	Kind      string   `json:"kind"`
	Tags      []uint16 `json:"tags,omitempty"`
	Part      string   `json:"part,omitempty"`
	PartWidth int      `json:"part_width,omitempty"`
	Keys      []string `json:"keys,omitempty"`
	Widths    []int    `json:"widths,omitempty"`
	Silk      string   `json:"silk,omitempty"`
}
