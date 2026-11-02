//ff:func feature=decl type=encoder control=sequence
//ff:what Extractor.langSlot — the language partition key; empty for an Edge without language
package decl

import "github.com/geul-org/silk/pkg/grammar"

func (x *Extractor) langSlot(lang []byte) []byte {
	return grammar.AppendSlot(nil, grammar.Val{B: lang}, lang != nil, x.d.PartWidth)
}
