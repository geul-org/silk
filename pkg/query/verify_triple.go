//ff:func feature=query type=matcher control=sequence
//ff:what verifyTriple — decode a hit with gogeul core: one TRIPLE record whose Predicate and Object equal the condition; its Edge Id
package query

import (
	"bytes"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// verifyTriple is the exact check after the byte scan.
func verifyTriple(rec []byte, t Triple) (uint64, error) {
	ws, err := core.ReadWords(rec)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	r, n, err := core.DecodeRecord(ws)
	if err != nil || n != len(ws) || r.Kind != core.KindEdge || r.Codec != core.EdgeTriple || len(r.Issuer) > 0 {
		return 0, fmt.Errorf("%w: not one TRIPLE record", ErrVerify)
	}
	fs, err := fields(r.Payload, 4)
	if err != nil {
		return 0, err
	}
	if (t.Pred != nil && !bytes.Equal(core.PutWords(fs[1]), t.Pred)) || (t.Obj != nil && !bytes.Equal(core.PutWords(fs[3]), t.Obj)) {
		return 0, fmt.Errorf("%w: Predicate or Object differs", ErrVerify)
	}
	e, _, _ := core.DecodeRecord(fs[0])
	id, err := core.DecodeLocalRef(e)
	if err != nil {
		return 0, fmt.Errorf("%w: Edge Id: %v", ErrVerify, err)
	}
	return id, nil
}
