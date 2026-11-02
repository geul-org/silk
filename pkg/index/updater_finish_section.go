//ff:func feature=index type=builder control=sequence
//ff:what updater.finishSection — publish a section's new rows; a pending section with no rows is simply built
package index

func (u *updater) finishSection(s *Section, p string, r *Report) error {
	if p == "" {
		s.Pending = false
		return nil
	}
	rows, err := rowsIn(p)
	if err != nil {
		return err
	}
	r.Rows += rows
	return u.publish(s, p, rows)
}
