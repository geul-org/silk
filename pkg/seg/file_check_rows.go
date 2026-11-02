//ff:func feature=seg type=rule control=iteration dimension=1
//ff:what File.checkRows — partition i's rows ascend
package seg

import (
	"bytes"
	"fmt"
)

func (f *File) checkRows(i uint64) error {
	_, r := f.PartAt(i)
	for j := 1; j < r.Len(); j++ {
		if bytes.Compare(r.Row(j-1), r.Row(j)) > 0 {
			return fmt.Errorf("%w: %s rows out of order", ErrFormat, f.Path)
		}
	}
	return nil
}
