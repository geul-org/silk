//ff:func feature=query type=helper control=sequence
//ff:what subjectQID — the QID word of a record's Subject (field 1 of a text or fact Edge)
package query

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/index"
)

func subjectQID(ix *index.Index, d *grammar.Doc, rec grammar.Rec) uint32 {
	s, ok := ix.Grammar.Get(d.At(rec), grammar.Path{1})
	if !ok {
		return 0
	}
	return grammar.EntityID(s.B)
}
