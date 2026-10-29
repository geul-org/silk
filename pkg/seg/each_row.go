//ff:func feature=seg type=reader control=iteration dimension=1
//ff:what eachRow — every row of one group with the group's key
package seg

// eachRow drains a cursor into fn.
func eachRow(g Group, fn func(part, row []byte) error) error {
	c := &cursor{g: g}
	for {
		row, ok, err := c.next()
		if !ok {
			return err
		}
		if err := fn(g.Key, row); err != nil {
			return err
		}
	}
}
