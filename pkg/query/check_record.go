//ff:func feature=query type=matcher control=iteration dimension=1
//ff:what checkRecord — decode a candidate with gogeul core: one Edge of the condition's codec whose fields equal the condition; its Edge Id, or false when a field differs
package query

import (
	"bytes"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkRecord is the exact check; key slots only narrowed the candidates.
func checkRecord(rec []byte, c Cond) (uint64, bool, error) {
	ws, err := core.ReadWords(rec)
	if err != nil {
		return 0, false, fmt.Errorf("%w: %v", ErrVerify, err)
	}
	r, n, err := core.DecodeRecord(ws)
	if err != nil || n != len(ws) || r.Kind != core.Kind(c.Kind) || r.Codec != c.Codec || len(r.Issuer) > 0 {
		return 0, false, fmt.Errorf("%w: not one %s record", ErrVerify, c.Name)
	}
	fs, err := fields(r.Payload, maxField(c)+1)
	if err != nil {
		return 0, false, err
	}
	for _, f := range c.Fields {
		if v, ok := fieldValue(fs, f.Path); !ok || !bytes.Equal(v, f.Value) {
			return 0, false, nil
		}
	}
	e, _, _ := core.DecodeRecord(fs[0])
	id, err := core.DecodeLocalRef(e)
	if err != nil {
		return 0, false, fmt.Errorf("%w: Edge Id: %v", ErrVerify, err)
	}
	return id, true, nil
}
