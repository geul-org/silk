//ff:func feature=recmap type=encoder control=iteration dimension=1
//ff:what encodeBody — key table (key, reserved, start, count), document table (off, len, ID, version, reserved), off[] u32, little-endian
package recmap

import "encoding/binary"

// encodeBody lays out the three tables after the header (docs/recmap.md).
func encodeBody(m *Map) []byte {
	le := binary.LittleEndian
	b := make([]byte, 0, len(m.Keys)*keySize+len(m.Docs)*docSize+len(m.Offs)*offSize)
	for _, k := range m.Keys {
		b = append(le.AppendUint16(b, k.Key), 0, 0, 0, 0, 0, 0)
		b = le.AppendUint64(le.AppendUint64(b, k.Start), k.Count)
	}
	for _, d := range m.Docs {
		b = le.AppendUint64(le.AppendUint32(le.AppendUint32(b, d.Off), d.Len), d.ID)
		b = le.AppendUint32(le.AppendUint32(b, d.Version), 0)
	}
	for _, o := range m.Offs {
		b = le.AppendUint32(b, o)
	}
	return b
}
