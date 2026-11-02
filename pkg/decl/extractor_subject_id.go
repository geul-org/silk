//ff:func feature=decl type=helper control=sequence
//ff:what Extractor.subjectID — the QID word of a record's Subject
package decl

import "github.com/geul-org/silk/pkg/grammar"

func (x *Extractor) subjectID(doc *grammar.Doc, r grammar.Rec) uint32 {
	s, ok := x.g.Get(doc.At(r), x.sub)
	if !ok {
		return 0
	}
	return grammar.EntityID(s.B)
}
