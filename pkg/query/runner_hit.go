//ff:func feature=query type=builder control=selection
//ff:what runner.hit — a live row as a hit, read by its section kind
package query

import (
	"encoding/binary"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/seg"
)

func (r *runner) hit(p *Plan, f *seg.File, row []byte) (Hit, bool) {
	seq := f.Doc(row)
	if !r.ix.Live.Get(seq) || !bitsMatch(p, row) {
		return Hit{}, false
	}
	h := Hit{Doc: seq, Key: row}
	switch p.Decl.Kind {
	case decl.KindIndex:
		h.Off, h.HasOff = binary.BigEndian.Uint32(row[f.DocOff+4:]), true
	case decl.KindText:
		h.QID = binary.BigEndian.Uint32(row[decl.TextWidth+4:])
	case decl.KindComposite:
		h.QID = binary.BigEndian.Uint32(row[decl.TextWidth+4:])
	case decl.KindSilk:
		h.Loose = !anyExact(p.Pats, row)
	}
	return h, true
}
