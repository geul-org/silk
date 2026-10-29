//ff:func feature=index type=reader control=sequence
//ff:what Index.Record — the bytes of the record at an offset of a document, read from its data pack and framed by its TAG WORD
package index

import (
	"fmt"

	"github.com/geul-org/silk/pkg/spec"
)

// Record reads record off of document doc.
func (ix *Index) Record(doc, off uint32) ([]byte, error) {
	if int(doc) >= len(ix.Docs) || off >= ix.Docs[doc].Len {
		return nil, fmt.Errorf("%w: no record at %d of document %d", ErrFormat, off, doc)
	}
	d := ix.Docs[doc]
	f, err := ix.pack(d.Pack)
	if err != nil {
		return nil, err
	}
	head := make([]byte, min(6, d.Len-off))
	if _, err := f.ReadAt(head, int64(d.Off)+int64(off)); err != nil {
		return nil, err
	}
	n := spec.RecLen(head)
	if n <= 0 || uint64(off)+uint64(n) > uint64(d.Len) {
		return nil, fmt.Errorf("%w: no record at %d of document %d", ErrFormat, off, doc)
	}
	rec := make([]byte, n)
	_, err = f.ReadAt(rec, int64(d.Off)+int64(off))
	return rec, err
}
