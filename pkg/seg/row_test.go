//ff:func feature=seg type=helper control=sequence
//ff:what row — a test row: key bytes then document number
package seg

import "encoding/binary"

func row(a, b byte, doc uint32) []byte {
	return binary.BigEndian.AppendUint32([]byte{a, b}, doc)
}
