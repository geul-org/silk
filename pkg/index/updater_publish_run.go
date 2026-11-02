//ff:func feature=index type=builder control=sequence
//ff:what updater.publishRun — a new sorted run: the new rows merged with the append array (the flush at the limit)
package index

func (u *updater) publishRun(s *Section, tmp string, rows uint64) error {
	if len(s.Appends) > 0 {
		merged, n, err := u.withAppends(s, tmp)
		if err != nil {
			return err
		}
		tmp, rows = merged, n
		s.Appends = nil
	}
	ref, err := u.place(tmp, rows)
	s.Runs = append(s.Runs, ref)
	s.Pending = false
	return err
}
