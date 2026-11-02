//ff:func feature=grammar type=decoder control=sequence
//ff:what Target — a record's target word: Kind and Codec of its TAG WORD, a Reference's Quantifier cleared
package grammar

import "encoding/binary"

func Target(rec []byte) uint16 {
	t := binary.BigEndian.Uint16(rec) & 0x187F
	if t>>11 == kindReference {
		t &^= 7
	}
	return t
}
