//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what Extractor.textRows — one row per text Edge: first 8 text bytes, document number, QID word
package decl

import (
	"encoding/binary"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/seg"
)

func (x *Extractor) textRows(doc *grammar.Doc, seq uint32, out *seg.Buffer) {
	var row [TextWidth + 8]byte
	for _, r := range doc.Recs {
		if !x.tags[r.Tag] {
			continue
		}
		lang, text, ok := doc.Text(x.g, doc.At(r))
		if !ok {
			continue
		}
		clear(row[:TextWidth])
		copy(row[:TextWidth], text)
		binary.BigEndian.PutUint32(row[TextWidth:], seq)
		binary.BigEndian.PutUint32(row[TextWidth+4:], x.subjectID(doc, r))
		out.Add(x.langSlot(lang), row[:])
	}
}
