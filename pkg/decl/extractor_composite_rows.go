//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what Extractor.compositeRows — one row per text Edge: first 8 text bytes, meaning 32 bits, QID word, document number
package decl

import (
	"encoding/binary"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/seg"
)

func (x *Extractor) compositeRows(doc *grammar.Doc, seq uint32, out *seg.Buffer) {
	bits := map[string][]byte{}
	for _, e := range x.def.Entities(x.g, doc) {
		bits[string(e.Subject)] = x.def.Single(e)
	}
	var row [TextWidth + 12]byte
	for _, r := range doc.Recs {
		if !x.tags[r.Tag] {
			continue
		}
		lang, text, ok := doc.Text(x.g, doc.At(r))
		if !ok {
			continue
		}
		clear(row[:TextWidth+4])
		copy(row[:TextWidth], text)
		s, _ := x.g.Get(doc.At(r), x.sub)
		copy(row[TextWidth:TextWidth+4], bits[string(s.B)])
		binary.BigEndian.PutUint32(row[TextWidth+4:], grammar.EntityID(s.B))
		binary.BigEndian.PutUint32(row[TextWidth+8:], seq)
		out.Add(x.langSlot(lang), row[:])
	}
}
