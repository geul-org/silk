//ff:func feature=decl type=encoder control=iteration dimension=1
//ff:what Extractor.indexRow — the partition slot and key slots of a record
package decl

import "github.com/geul-org/silk/pkg/grammar"

func (x *Extractor) indexRow(doc *grammar.Doc, rec, part, row []byte) ([]byte, []byte) {
	if x.part != nil {
		v, ok := doc.Get(x.g, rec, x.part.Path)
		part = grammar.AppendSlot(part, v, ok, x.d.PartWidth)
	}
	for i, k := range x.keys {
		v, ok := doc.Get(x.g, rec, k.Path)
		row = grammar.AppendSlot(row, v, ok, x.d.Widths[i])
	}
	return part, row
}
