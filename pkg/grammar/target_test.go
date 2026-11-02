//ff:func feature=grammar type=scanner control=sequence
//ff:what TestTarget — Kind and Codec of a TAG WORD, Length gone, a Reference's Quantifier cleared
package grammar

import "testing"

func TestTarget(t *testing.T) {
	if Target([]byte{0x16, 0x8a}) != 0x100A || Target([]byte{0x00, 0xA4}) != 0x0020 || Target([]byte{0x0A, 0xCB}) != 0x084B {
		t.Fatal("target")
	}
}
