//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what Extractor.silkRows — each entity's rows: bits then document number, and document number then bits
package decl

import (
	"encoding/binary"
	"slices"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/seg"
)

func (x *Extractor) silkRows(doc *grammar.Doc, seq uint32, out []*seg.Buffer) {
	for _, e := range x.def.Entities(x.g, doc) {
		for _, bits := range x.def.Rows(e) {
			s := binary.BigEndian.AppendUint32(nil, seq)
			out[0].Add(nil, append(slices.Clone(bits), s...))
			out[1].Add(nil, append(s, bits...))
		}
	}
}
