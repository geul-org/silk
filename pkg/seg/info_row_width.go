//ff:func feature=seg type=helper control=sequence
//ff:what Info.RowWidth — bytes of a partition row (sort slot and sequence) and of an other row (partition slot in front)
package seg

// RowWidth returns (partition row, other row) widths.
func (i Info) RowWidth() (int, int) {
	return i.SortWidth + seqSize, i.PartWidth + i.SortWidth + seqSize
}
