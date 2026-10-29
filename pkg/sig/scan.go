//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what Scan — the rows of a table a pattern accepts, scanned by several goroutines in contiguous chunks, in row order
package sig

import "sync"

// Scan runs scanRows over threads chunks (threads ≤ 0 means 8).
func Scan(t *Table, p Pattern, threads int) []uint32 {
	n := t.Len()
	if threads <= 0 {
		threads = 8
	}
	threads = max(1, min(threads, n))
	parts := make([][]uint32, threads)
	var wg sync.WaitGroup
	for i := range threads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			parts[i] = scanRows(t, p, n*i/threads, n*(i+1)/threads)
		}()
	}
	wg.Wait()
	var out []uint32
	for _, part := range parts {
		out = append(out, part...)
	}
	return out
}
