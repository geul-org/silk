//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what File.emitRange — read rows [lo, hi) of a region in chunks and emit those that match (all when cmp is false)
package seg

// emitRange reads chunkRows rows per ReadAt.
func (f *File) emitRange(rg region, lo, hi uint64, prefix []byte, at int, cmp bool, emit func(doc, off uint32) error) error {
	for lo < hi {
		k := min(uint64(chunkRows), hi-lo)
		b, err := f.readRows(rg, lo, k)
		if err != nil {
			return err
		}
		if err := emitChunk(b, rg.w, prefix, at, cmp, emit); err != nil {
			return err
		}
		lo += k
	}
	return nil
}
