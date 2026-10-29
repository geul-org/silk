//ff:func feature=index type=encoder control=iteration dimension=1
//ff:what encodeDocs — a document table file: header (first sequence, count, pair count, body SHA-256), 40-byte entries, then (TAG, count) pairs
package index

import (
	"crypto/sha256"
	"encoding/binary"
)

// encodeDocs lays out docs/index.md "Document table".
func encodeDocs(first uint32, ds []DocEntry) []byte {
	be := binary.BigEndian
	var body, pairs []byte
	np := uint32(0)
	for _, d := range ds {
		body = be.AppendUint64(be.AppendUint32(be.AppendUint32(body, d.Pack), d.Issuer), d.Off)
		body = be.AppendUint64(be.AppendUint32(be.AppendUint32(body, d.Len), d.Version), d.ID)
		body = be.AppendUint16(be.AppendUint16(be.AppendUint32(body, np), uint16(len(d.Counts))), 0)
		pairs = appendPairs(pairs, d.Counts)
		np += uint32(len(d.Counts))
	}
	body = append(body, pairs...)
	out := be.AppendUint16(append([]byte{}, docsMagic...), fileVersion)
	out = be.AppendUint32(be.AppendUint32(be.AppendUint32(append(out, 0, 0), first), uint32(len(ds))), np)
	s := sha256.Sum256(body)
	return append(append(out, s[:]...), body...)
}
