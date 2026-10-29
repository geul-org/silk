//ff:func feature=seg type=reader control=sequence
//ff:what File.selectRegion — rows of a region whose bytes at `at` start with prefix: binary search in a sorted run (prefix at 0), else a scan; the rows examined
package seg

// selectRegion emits each matching row's document and offset.
func (f *File) selectRegion(rg region, prefix []byte, at int, emit func(doc, off uint32) error) (uint64, error) {
	lo, hi, cmp := uint64(0), rg.n, true
	if f.Info.Sorted && at == 0 && len(prefix) > 0 {
		var err error
		if lo, err = f.bound(rg, prefix, false); err != nil {
			return 0, err
		}
		if hi, err = f.bound(rg, prefix, true); err != nil {
			return 0, err
		}
		cmp = false
	}
	return hi - lo, f.emitRange(rg, lo, hi, prefix, at, cmp, emit)
}
