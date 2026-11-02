//ff:func feature=query type=rule control=sequence
//ff:what flagged — a rest field's slot in an INDEX ends in its whole-or-cut byte; text keys are plain bytes
package query

import (
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

func flagged(d decl.Decl, l grammar.Leaf) bool {
	return d.Kind == decl.KindIndex && l.Field.Width == grammar.RestWidth
}
