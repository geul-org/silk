//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what fields — the first n embedded records of an Edge payload, each as its words (core.DecodeRecord framing)
package query

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// fields splits ws into n records; fewer is ErrVerify.
func fields(ws []core.Word, n int) ([][]core.Word, error) {
	out := make([][]core.Word, 0, n)
	for len(out) < n {
		_, m, err := core.DecodeRecord(ws)
		if err != nil {
			return nil, fmt.Errorf("%w: field %d: %v", ErrVerify, len(out), err)
		}
		out = append(out, ws[:m])
		ws = ws[m:]
	}
	return out, nil
}
