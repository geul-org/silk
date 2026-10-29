//ff:func feature=recmap type=reader control=sequence
//ff:what Read — a map file of size bytes from r: header, key table and document table checked; off[] is left on r
package recmap

import (
	"fmt"
	"io"
)

// Read parses the header and the two small tables. The file size must be
// exactly header + tables + off[], so counts cannot ask for more than exists.
func Read(r io.ReaderAt, size int64) (*File, error) {
	hb := make([]byte, headerSize)
	if _, err := r.ReadAt(hb, 0); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFormat, err)
	}
	h, err := parseHeader(hb)
	if err != nil {
		return nil, err
	}
	if uint64(size) != bodyEnd(h) {
		return nil, fmt.Errorf("%w: %d bytes, header says %d", ErrFormat, size, bodyEnd(h))
	}
	tb := make([]byte, uint64(h.Keys)*keySize+h.Docs*docSize)
	if _, err := r.ReadAt(tb, headerSize); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFormat, err)
	}
	keys, err := parseKeys(tb[:uint64(h.Keys)*keySize], h.Records)
	if err != nil {
		return nil, err
	}
	docs, err := parseDocs(tb[uint64(h.Keys)*keySize:])
	if err != nil {
		return nil, err
	}
	return &File{Header: h, Keys: keys, Docs: docs, r: r}, nil
}
