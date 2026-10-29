//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what File.Each — every row of a segment with its partition key slot, partition by partition, then the other partition
package seg

// Each calls fn(part, row) for each row (row without the partition slot);
// both slices are valid only during the call.
func (f *File) Each(fn func(part, row []byte) error) error {
	gs, err := f.Groups()
	if err != nil {
		return err
	}
	for _, g := range gs {
		if err := eachRow(g, fn); err != nil {
			return err
		}
	}
	return nil
}
