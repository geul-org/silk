//ff:func feature=query type=matcher control=iteration dimension=1
//ff:what checkChunk — read and check each candidate of a chunk, filling its Edge Id and match flag
package query

import (
	"fmt"

	"github.com/park-jun-woo/silk/pkg/index"
)

// checkChunk is one goroutine's share of checkAll.
func checkChunk(ix *index.Index, cs []candidate, c Cond, ids []uint64, ok []bool) error {
	for i, k := range cs {
		rec, err := ix.Record(k.doc, k.off)
		if err != nil {
			return err
		}
		if ids[i], ok[i], err = checkRecord(rec, c); err != nil {
			return fmt.Errorf("%s offset %d: %w", ix.Path(k.doc), k.off, err)
		}
	}
	return nil
}
