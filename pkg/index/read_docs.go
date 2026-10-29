//ff:func feature=index type=reader control=sequence
//ff:what readDocs — a document table file: header and body SHA-256 checked, entries and their record counts decoded
package index

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
)

// readDocs returns the first sequence number and the entries of path.
func readDocs(path string) (uint32, []DocEntry, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, nil, err
	}
	be := binary.BigEndian
	if len(b) < docsHead || string(b[:8]) != docsMagic || be.Uint16(b[8:]) != fileVersion {
		return 0, nil, fmt.Errorf("%w: %s", ErrFormat, path)
	}
	first, n, np := be.Uint32(b[12:]), uint64(be.Uint32(b[16:])), uint64(be.Uint32(b[20:]))
	body := b[docsHead:]
	s := sha256.Sum256(body)
	if uint64(len(body)) != n*docSize+np*pairSize || !bytes.Equal(s[:], b[24:56]) {
		return 0, nil, fmt.Errorf("%w: %s body", ErrFormat, path)
	}
	ds, err := parseDocs(body[:n*docSize], body[n*docSize:])
	if err != nil {
		return 0, nil, fmt.Errorf("%s: %w", path, err)
	}
	return first, ds, nil
}
