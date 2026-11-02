//ff:type feature=grammar type=model
//ff:what Grammar — what SILK keeps of `0x0000`: fields by target word, target words by name, odd members of byte-string pairs, language short codes by lower-case BCP 47 tag
package grammar

type Grammar struct {
	Fields map[uint16][]Field  `json:"fields"`
	Names  map[string][]uint16 `json:"names"`
	Odd    map[uint16]bool     `json:"odd"`
	Langs  map[string]uint16   `json:"langs"`
}
