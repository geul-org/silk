//ff:func feature=recmap type=encoder control=sequence
//ff:what encodeHeader — the 176-byte header: magic, version, Issuer length, counts, generation, pack size, pack and manifest SHA-256, Issuer, body SHA-256
package recmap

import (
	"encoding/binary"
	"fmt"
)

// encodeHeader writes h; Issuer, Manifest and Pack must be well-formed hex.
func encodeHeader(h Header) ([]byte, error) {
	iss, err := issuerBytes(h.Issuer)
	if err != nil {
		return nil, err
	}
	pk, err := hashBytes(h.Pack)
	if err != nil {
		return nil, fmt.Errorf("pack: %w", err)
	}
	mf, err := hashBytes(h.Manifest)
	if err != nil {
		return nil, fmt.Errorf("manifest: %w", err)
	}
	le := binary.LittleEndian
	b := append(make([]byte, 0, headerSize), magic...)
	b = le.AppendUint16(le.AppendUint16(b, version), uint16(len(iss)))
	b = le.AppendUint64(le.AppendUint64(le.AppendUint32(b, h.Keys), h.Records), h.Docs)
	b = le.AppendUint64(le.AppendUint64(b, h.Generation), h.PackSize)
	b = append(append(b, pk[:]...), mf[:]...)
	b = append(append(b, iss...), make([]byte, maxIssuer-len(iss))...)
	return append(b, h.Sum[:]...), nil
}
