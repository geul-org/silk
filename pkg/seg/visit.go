//ff:func feature=seg type=scanner control=sequence
//ff:what visit — hand one live row to fn
package seg

func visit(c *cursor, hk Hooks, fn func(part, row []byte) error) error {
	if hk.Live != nil && !hk.Live(c.f.Doc(c.row())) {
		return nil
	}
	return fn(c.key, c.row())
}
