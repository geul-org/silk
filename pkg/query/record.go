//ff:func feature=query type=reader control=sequence
//ff:what Record — the record bytes of a hit; nil when it has no offset
package query

import (
	"context"
	"fmt"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/index"
)

func Record(ctx context.Context, ix *index.Index, h Hit) ([]byte, error) {
	if !h.HasOff {
		return nil, nil
	}
	b, err := ix.Bytes(ctx, h.Doc)
	if err != nil {
		return nil, err
	}
	if int(h.Off) >= len(b) {
		return nil, fmt.Errorf("%w: offset %d past document", index.ErrFormat, h.Off)
	}
	return b[h.Off : int(h.Off)+grammar.RecLen(b[h.Off:])], nil
}
