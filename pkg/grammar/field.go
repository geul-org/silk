//ff:type feature=grammar type=model
//ff:what Field — one declared payload field: name, width in words (0 record, 65535 rest), allowed CODEC targets
package grammar

type Field struct {
	Name   string   `json:"name"`
	Width  int      `json:"width"`
	Codecs []uint16 `json:"codecs,omitempty"`
}
