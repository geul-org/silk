//ff:type feature=grammar type=model
//ff:what Rec — one top-level record of a document: offset, length and target word
package grammar

type Rec struct {
	Off int
	Len int
	Tag uint16
}
