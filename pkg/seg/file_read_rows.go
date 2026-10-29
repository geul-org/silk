//ff:func feature=seg type=reader control=sequence
//ff:what File.readRows — rows [start, start+count) of a region in one ReadAt
package seg

import "fmt"

// readRows reads count rows from row start of rg.
func (f *File) readRows(rg region, start, count uint64) ([]byte, error) {
	b := make([]byte, count*uint64(rg.w))
	if _, err := f.r.ReadAt(b, rg.off+int64(start)*int64(rg.w)); err != nil {
		return nil, fmt.Errorf("%w: %s: %v", ErrFormat, f.Path, err)
	}
	return b, nil
}
