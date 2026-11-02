//ff:func feature=index type=builder control=sequence
//ff:what updater.publish — give a new sorted run its place: the append array while it fits, else a run together with the appends
package index

func (u *updater) publish(s *Section, tmp string, rows uint64) error {
	if s.Pending || rows+rowsOf(s.Appends) >= u.o.Policy.AppendRows {
		return u.publishRun(s, tmp, rows)
	}
	ref, err := u.place(tmp, rows)
	s.Appends = append(s.Appends, ref)
	return err
}
