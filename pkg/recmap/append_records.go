//ff:func feature=recmap type=builder control=iteration dimension=1
//ff:what appendRecords — validate one document with core.SplitDocument, then append each top-level record's key and pack offset
package recmap

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// appendRecords frames the document b, which starts at pack offset base.
// SplitDocument checks DOC-*, HEAD/BODY placement and END; DecodeRecord gives
// each record's Kind, Codec and word count.
func appendRecords(keys []uint16, offs []uint32, b []byte, base uint32) ([]uint16, []uint32, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrDocument, err)
	}
	if _, err := core.SplitDocument(ws); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrDocument, err)
	}
	for i := 0; i < len(ws); {
		r, n, _ := core.DecodeRecord(ws[i:])
		keys = append(keys, uint16(core.TargetWord(r.Kind, r.Codec)))
		offs = append(offs, base+uint32(2*i))
		i += n
	}
	return keys, offs, nil
}
