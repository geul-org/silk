//ff:func feature=grammar type=helper control=sequence
//ff:what grammarTarget — the target word of encoded bytes
package grammar

func grammarTarget(b []byte) uint16 {
	return Target(b)
}
