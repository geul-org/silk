//ff:func feature=seg type=helper control=sequence
//ff:what cursor.row — the current row
package seg

func (c *cursor) row() []byte {
	return c.reg.Row(c.i)
}
