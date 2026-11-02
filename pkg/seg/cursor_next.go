//ff:func feature=seg type=scanner control=sequence
//ff:what cursor.next — step one row; false at the end
package seg

func (c *cursor) next() bool {
	c.i++
	if c.i < c.reg.Len() {
		return true
	}
	return c.load()
}
