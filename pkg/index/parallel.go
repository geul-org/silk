//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what parallel — run fn(0..n-1) on at most threads goroutines; the first error
package index

import (
	"cmp"
	"sync"
)

func parallel(n, threads int, fn func(int) error) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var first error
	sem := make(chan struct{}, threads)
	for i := range n {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer func() { <-sem; wg.Done() }()
			if err := fn(i); err != nil {
				mu.Lock()
				first = cmp.Or(first, err)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return first
}
