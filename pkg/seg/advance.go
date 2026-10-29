//ff:func feature=seg type=reader control=sequence
//ff:what advance — a cursor's next row, nil at its end
package seg

// advance wraps cursor.next for the merge loop.
func advance(c *cursor) ([]byte, error) {
	row, ok, err := c.next()
	if !ok {
		return nil, err
	}
	return row, nil
}
