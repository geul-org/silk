//ff:func feature=scan type=scanner control=iteration dimension=1
//ff:what Records — scan the records at ascending pack offsets with a Matcher over several goroutines; hits come back in offset order
package scan

import (
	"io"
	"sync"
)

// Records splits offs into o.Threads contiguous chunks, each scanned by its
// own goroutine with its own window over data (a data pack of size bytes).
func Records(data io.ReaderAt, size uint64, offs []uint32, m Matcher, o Options) ([]Hit, error) {
	th, win := o.resolve(len(offs))
	parts := make([][]Hit, th)
	errs := make([]error, th)
	var wg sync.WaitGroup
	for i := range th {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w := &window{data: data, size: size, min: win}
			parts[i], errs[i] = scanChunk(w, offs[len(offs)*i/th:len(offs)*(i+1)/th], m)
		}()
	}
	wg.Wait()
	return join(parts, errs)
}
