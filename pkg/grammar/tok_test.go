//ff:func feature=grammar type=helper control=sequence
//ff:what tok — a Reference token's record
package grammar

func tok(s string) []byte {
	b, _ := Ref(s)
	return b
}
