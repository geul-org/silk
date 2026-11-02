//ff:type feature=query type=model
//ff:what Pred — one field predicate: resolved path, the byte strings it may equal, and whether it asks for a prefix
package query

import "github.com/geul-org/silk/pkg/grammar"

type Pred struct {
	Text   string
	Leaf   grammar.Leaf
	Vals   [][]byte
	Prefix bool
}
