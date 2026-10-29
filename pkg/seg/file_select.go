//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what File.Select — the rows of a selection: one partition (or the other partition's slice for that key) or every partition and the other one; rows examined
package seg

// Select emits the document and offset of each selected row. With a
// partition key, a key that has no partition of its own is looked up in the
// other partition (rows sorted by partition slot, then sort slot).
func (f *File) Select(sel Sel, emit func(doc, off uint32) error) (uint64, error) {
	if sel.Part != nil {
		if i, ok := f.Find(sel.Part); ok {
			return f.selectRegion(f.partRegion(i), sel.Sort, 0, emit)
		}
		return f.selectRegion(f.otherRegion(), append(append([]byte{}, sel.Part...), sel.Sort...), 0, emit)
	}
	total := uint64(0)
	for i := range f.Parts {
		n, err := f.selectRegion(f.partRegion(i), sel.Sort, 0, emit)
		if err != nil {
			return total, err
		}
		total += n
	}
	n, err := f.selectRegion(f.otherRegion(), sel.Sort, f.Info.PartWidth, emit)
	return total + n, err
}
