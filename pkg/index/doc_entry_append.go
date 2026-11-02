//ff:func feature=index type=encoder control=sequence
//ff:what DocEntry.append — the 40-byte table entry
package index

import "encoding/binary"

func (e DocEntry) append(b []byte) []byte {
	b = binary.BigEndian.AppendUint32(b, e.Pack)
	b = binary.BigEndian.AppendUint32(b, e.Issuer)
	b = binary.BigEndian.AppendUint64(b, e.Off)
	b = binary.BigEndian.AppendUint64(b, e.Len)
	b = binary.BigEndian.AppendUint64(b, e.ID)
	b = binary.BigEndian.AppendUint32(b, e.Version)
	return append(b, e.Enc, 0, 0, 0)
}
