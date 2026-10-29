//ff:func feature=seg type=encoder control=sequence
//ff:what AppendSeq — append a document sequence number and record offset to a row (big-endian, so rows sort by them after the key)
package seg

import "encoding/binary"

// AppendSeq finishes a row.
func AppendSeq(row []byte, doc, off uint32) []byte {
	return binary.BigEndian.AppendUint32(binary.BigEndian.AppendUint32(row, doc), off)
}
