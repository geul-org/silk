//ff:func feature=index type=encoder control=sequence
//ff:what identity — the 16-byte key of a document version
package index

import "encoding/binary"

func identity(e DocEntry) []byte {
	b := binary.BigEndian.AppendUint32(make([]byte, 0, pathKey+4), e.Issuer)
	b = binary.BigEndian.AppendUint64(b, e.ID)
	return binary.BigEndian.AppendUint32(b, e.Version)
}
