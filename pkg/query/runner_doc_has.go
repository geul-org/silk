//ff:func feature=query type=rule control=iteration dimension=1
//ff:what runner.docHas — a document's rows of the document-order bundle meet a pattern; a row of mere candidates goes to the document, unless only rows may be read
package query

import (
	"encoding/binary"

	"github.com/geul-org/silk/pkg/index"
)

func (r *runner) docHas(p *Plan, sec *index.Sec, seq uint32, rowsOnly bool) (bool, error) {
	key := binary.BigEndian.AppendUint32(nil, seq)
	loose := false
	for _, f := range sec.Files {
		exact, l := rowsMeet(p.Pats, f.All().Prefix(key))
		if exact {
			return true, nil
		}
		loose = loose || l
	}
	if !loose || rowsOnly {
		return loose, nil
	}
	yes, unknown, err := r.holds(p, seq)
	return yes || unknown, err
}
