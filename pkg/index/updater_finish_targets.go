//ff:func feature=index type=builder control=iteration dimension=1
//ff:what updater.finishTargets — each spill down to one run, placed in section order
package index

func (u *updater) finishTargets(ts []*target, r *Report) error {
	for _, t := range ts {
		if err := u.finishTarget(t, r); err != nil {
			return err
		}
	}
	return nil
}
