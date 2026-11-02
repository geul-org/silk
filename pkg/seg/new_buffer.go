//ff:func feature=seg type=builder control=sequence
//ff:what NewBuffer — an empty buffer for rows of this shape
package seg

func NewBuffer(i Info) *Buffer {
	return &Buffer{pw: i.PartWidth, rw: i.RowWidth}
}
