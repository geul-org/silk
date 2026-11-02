//ff:type feature=sig type=model
//ff:what Def — a parsed definition: name, width in bits, common fields, when blocks and the hash of its text
package sig

type Def struct {
	Name   string
	Bits   int
	Common []*Field
	Whens  []*When
	Hash   [32]byte
	Preds  map[string]bool
}
