//ff:func feature=seg type=writer control=sequence
//ff:what emit — write the cursor's row unless its document is dead; whether it was kept
package seg

func emit(w *Writer, c *cursor, hk Hooks) (bool, error) {
	if hk.Live != nil && !hk.Live(c.f.Doc(c.row())) {
		return false, nil
	}
	return true, w.Add(c.key, c.row())
}
