//ff:func feature=index type=builder control=iteration dimension=1
//ff:what updater.finishTarget — one declaration's bundles
package index

func (u *updater) finishTarget(t *target, r *Report) error {
	for i, sp := range t.sp {
		p, err := sp.finish(fanIn(u.o.Limit))
		if err != nil {
			return err
		}
		if err := u.finishSection(t.sec[i], p, r); err != nil {
			return err
		}
	}
	return nil
}
