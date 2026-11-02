//ff:type feature=grammar type=model
//ff:what Val — a field's content: a record's bytes, a raw field's words, or a rest field's bytes without padding
package grammar

type Val struct {
	B    []byte
	Raw  bool
	Rest bool
}
