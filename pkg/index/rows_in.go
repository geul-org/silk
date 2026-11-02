//ff:func feature=index type=reader control=sequence
//ff:what rowsIn — the rows of a segment
package index

import "github.com/geul-org/silk/pkg/seg"

func rowsIn(path string) (uint64, error) {
	f, err := seg.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Rows, nil
}
