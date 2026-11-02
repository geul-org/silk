//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what Extractor.indexRows — one row per record of the TAGs: slots, document number, offset
package decl

import (
	"encoding/binary"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/seg"
)

func (x *Extractor) indexRows(doc *grammar.Doc, seq uint32, out *seg.Buffer) {
	var part, row []byte
	for _, r := range doc.Recs {
		if !x.tags[r.Tag] {
			continue
		}
		part, row = x.indexRow(doc, doc.At(r), part[:0], row[:0])
		row = binary.BigEndian.AppendUint32(binary.BigEndian.AppendUint32(row, seq), uint32(r.Off))
		out.Add(part, row)
	}
}
