//ff:func feature=seg type=scanner control=sequence
//ff:what cursor.load — move to the next partition; false at the end
package seg

func (c *cursor) load() bool {
	if c.part >= c.f.Parts {
		return false
	}
	c.key, c.reg = c.f.PartAt(c.part)
	c.part++
	c.i = 0
	return true
}
