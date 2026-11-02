//ff:func feature=decl type=builder control=selection
//ff:what Extractor.Rows — the rows of one document number, into one buffer per bundle
package decl

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/seg"
)

func (x *Extractor) Rows(doc *grammar.Doc, seq uint32, out []*seg.Buffer) {
	switch x.d.Kind {
	case KindIndex:
		x.indexRows(doc, seq, out[0])
	case KindText:
		x.textRows(doc, seq, out[0])
	case KindComposite:
		x.compositeRows(doc, seq, out[0])
	case KindSilk:
		x.silkRows(doc, seq, out)
	}
}
