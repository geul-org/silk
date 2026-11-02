//ff:type feature=decl type=model
//ff:what Extractor — one declaration compiled to make rows from documents
package decl

import (
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/sig"
)

type Extractor struct {
	d    Decl
	g    *grammar.Grammar
	tags map[uint16]bool
	part *grammar.Leaf
	keys []grammar.Leaf
	def  *sig.Def
	sub  grammar.Path
}
