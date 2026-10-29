//ff:func feature=recmap type=builder control=iteration dimension=1
//ff:what Build — the record map of documents stored raw in a data pack: every top-level record's offset sorted by (Kind·Codec, offset)
package recmap

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// Build reads each document of docs (ascending, non-overlapping, inside size)
// from data and frames it with gogeul core. The header gets the counts only;
// the caller names the Issuer, generation, manifest and pack.
func Build(data io.ReaderAt, size uint64, docs []archive.Doc) (*Map, error) {
	if err := checkDocs(docs, size); err != nil {
		return nil, err
	}
	var keys []uint16
	var offs []uint32
	var buf []byte
	ents := make([]DocEntry, 0, len(docs))
	for _, d := range docs {
		buf = grow(buf, int(d.Len))
		if _, err := data.ReadAt(buf, int64(d.Off)); err != nil {
			return nil, fmt.Errorf("document 0x%016x at %d: %w", d.ID, d.Off, err)
		}
		var err error
		if keys, offs, err = appendRecords(keys, offs, buf, uint32(d.Off)); err != nil {
			return nil, fmt.Errorf("document 0x%016x: %w", d.ID, err)
		}
		ents = append(ents, DocEntry{Off: uint32(d.Off), Len: uint32(d.Len), ID: d.ID, Version: d.Version})
	}
	ranges, sorted := sortByKey(keys, offs)
	h := Header{PackSize: size, Keys: uint32(len(ranges)), Records: uint64(len(sorted)), Docs: uint64(len(ents))}
	return &Map{Header: h, Keys: ranges, Docs: ents, Offs: sorted}, nil
}
