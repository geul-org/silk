//ff:func feature=grammar type=helper control=sequence
//ff:what Val.B2 — a test check: the value is a whole 2-word record
package grammar

func (v Val) B2() bool {
	return len(v.B) == 4
}
