//ff:func feature=seg type=reader control=sequence
//ff:what cursor.next — the next row (valid until the cursor reads its next chunk); false at the end
package seg

// next reads chunkRows rows at a time.
func (c *cursor) next() ([]byte, bool, error) {
	if c.i >= len(c.buf) {
		if c.read >= c.g.rg.n {
			return nil, false, nil
		}
		k := min(uint64(chunkRows), c.g.rg.n-c.read)
		b, err := c.g.f.readRows(c.g.rg, c.read, k)
		if err != nil {
			return nil, false, err
		}
		c.buf, c.i, c.read = b, 0, c.read+k
	}
	row := c.buf[c.i+c.g.strip : c.i+c.g.rg.w]
	c.i += c.g.rg.w
	return row, true, nil
}
