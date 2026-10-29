//ff:func feature=index type=writer control=sequence
//ff:what saveRound — write checkpoints, commit when merges were applied, collect unreferenced files; ErrPaused when checkpoints remain
package index

// saveRound ends a compaction round.
func saveRound(dir string, m *Manifest, cps []checkpoint, applied int, rep *Report) error {
	n, err := writeCheckpoints(dir, cps)
	rep.Written += n
	if err != nil {
		return err
	}
	if applied > 0 {
		n, err := commit(dir, m)
		rep.Written += n
		if err != nil {
			return err
		}
		rep.Generation = m.Generation
		collect(dir, *m)
	}
	if len(cps) > 0 {
		rep.Paused = true
		return ErrPaused
	}
	return nil
}
