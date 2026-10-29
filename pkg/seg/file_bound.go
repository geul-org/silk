//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what File.bound — binary search in a sorted region: the first row whose leading bytes are >= prefix (or > prefix when upper)
package seg

import "bytes"

// bound reads one row prefix per probe with ReadAt.
func (f *File) bound(rg region, prefix []byte, upper bool) (uint64, error) {
	lo, hi := uint64(0), rg.n
	b := make([]byte, len(prefix))
	for lo < hi {
		mid := lo + (hi-lo)/2
		if _, err := f.r.ReadAt(b, rg.off+int64(mid)*int64(rg.w)); err != nil {
			return 0, err
		}
		c := bytes.Compare(b, prefix)
		if c < 0 || (upper && c == 0) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo, nil
}
