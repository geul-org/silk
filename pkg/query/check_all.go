//ff:func feature=query type=matcher control=iteration dimension=1
//ff:what checkAll — the exact check of every candidate on up to `threads` goroutines; per candidate its Edge Id and whether it matched
package query

import (
	"sync"

	"github.com/geul-org/silk/pkg/index"
)

// checkAll splits cs into contiguous chunks; the first error wins.
func checkAll(ix *index.Index, cs []candidate, c Cond, threads int) ([]uint64, []bool, error) {
	ids, ok := make([]uint64, len(cs)), make([]bool, len(cs))
	th := max(1, min(threads, len(cs)))
	errs := make([]error, th)
	var wg sync.WaitGroup
	for i := range th {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lo, hi := len(cs)*i/th, len(cs)*(i+1)/th
			errs[i] = checkChunk(ix, cs[lo:hi], c, ids[lo:hi], ok[lo:hi])
		}()
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, nil, err
		}
	}
	return ids, ok, nil
}
