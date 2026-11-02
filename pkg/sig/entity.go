//ff:type feature=sig type=model
//ff:what Entity — the statements of one Subject in a document, deprecated ones left out
package sig

type Entity struct {
	Subject []byte
	Objects map[string][][]byte
}
