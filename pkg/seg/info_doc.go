//ff:func feature=seg type=decoder control=sequence
//ff:what Info.Doc — the document number of a row
package seg

import "encoding/binary"

func (i Info) Doc(row []byte) uint32 {
	return binary.BigEndian.Uint32(row[i.DocOff:])
}
