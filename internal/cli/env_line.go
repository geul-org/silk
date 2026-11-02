//ff:func feature=cli type=builder control=selection
//ff:what env.line — one hit as text; `~` marks a candidate, `?` a value that may hold but is unknown
package cli

import (
	"fmt"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/query"
)

func (e *env) line(ix *index.Index, mode string, h query.Hit) (string, error) {
	s := ix.Path(h.Doc)
	if h.Loose {
		s += " ~"
	}
	if h.Unknown {
		s += " ?"
	}
	switch mode {
	case query.ModeKeys:
		return fmt.Sprintf("%s %x qid=%08x", s, h.Key, h.QID), nil
	case query.ModeRecords:
		rec, err := query.Record(e.ctx, ix, h)
		return fmt.Sprintf("%s +%d %x", s, h.Off, rec), err
	}
	n, ok, err := query.EdgeID(e.ctx, ix, h)
	if ok {
		s += fmt.Sprintf(" #%d", n)
	}
	return s, err
}
