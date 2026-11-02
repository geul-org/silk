//ff:func feature=index type=builder control=sequence
//ff:what compactor.planSection — major when asked or stale past the ratio; else the fullest tier at its run count
package index

func (c *compactor) planSection(s *Section) (task, bool, error) {
	n := len(s.Runs) + len(s.Appends)
	if n == 0 {
		return task{}, false, nil
	}
	stale, err := staleShare(c.dir, *s, c.live)
	if err != nil {
		return task{}, false, err
	}
	if c.full && (n > 1 || stale > 0) || stale > c.o.Policy.StaleRatio {
		return task{sec: s, runs: s.Runs, appends: s.Appends}, true, nil
	}
	runs := tierRuns(s.Runs, c.o.Policy.RunsPerTier)
	return task{sec: s, runs: runs}, runs != nil, nil
}
