//ff:type feature=query type=model
//ff:what Hit — one answer: document, record offset when known, QID word, candidate row and whether it is only a candidate
package query

type Hit struct {
	Doc     uint32
	Off     uint32
	HasOff  bool
	QID     uint32
	Key     []byte
	Unknown bool
	Loose   bool
}
