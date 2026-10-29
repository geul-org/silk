//ff:func feature=recmap type=encoder control=sequence
//ff:what Map.Encode — the map file bytes: header (counts and body SHA-256 filled in) then key table, document table, off[]
package recmap

import "crypto/sha256"

// Encode sets m.Header's counts and Sum from m and returns the file bytes.
func (m *Map) Encode() ([]byte, error) {
	body := encodeBody(m)
	m.Header.Keys = uint32(len(m.Keys))
	m.Header.Records = uint64(len(m.Offs))
	m.Header.Docs = uint64(len(m.Docs))
	m.Header.Sum = sha256.Sum256(body)
	head, err := encodeHeader(m.Header)
	if err != nil {
		return nil, err
	}
	return append(head, body...), nil
}
