//ff:type feature=grammar type=model
//ff:what Doc — a framed document: its bytes, top-level records and lazily read REFERENCE-NODE bindings
package grammar

type Doc struct {
	B    []byte
	Recs []Rec
	refs map[string][]byte
}
