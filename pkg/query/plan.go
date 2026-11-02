//ff:type feature=query type=model
//ff:what Plan — one condition on one section: partitions, key prefixes and whether rows alone decide
package query

import (
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

type Plan struct {
	Cond    *Cond
	Name    string
	Decl    decl.Decl
	Parts   [][]byte
	Prefix  [][]byte
	Fixed   int
	Covered bool
	Pats    []sig.Pattern
	Def     *sig.Def
	Bits    []sig.Pattern
}
