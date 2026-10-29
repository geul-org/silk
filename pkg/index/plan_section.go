//ff:func feature=index type=rule control=sequence
//ff:what planSection — one section's due merge: major (more than MaxRuns segments, dead share above StaleRatio, or full), else flush (appended share above AppendRatio), else a tier merge
package index

// planSection applies the triggers in order of cost: a major merge also
// covers what a flush or a tier merge would do.
func planSection(i int, s Section, p Policy, full bool) (Task, bool) {
	apps := rowsOf(s.Appends)
	total := rowsOf(s.Runs) + apps
	if total == 0 {
		return Task{}, false
	}
	all := append(append([]SegRef{}, s.Runs...), s.Appends...)
	t := Task{Section: i, Name: s.Name}
	if (full && (len(all) > 1 || s.Dead > 0 || len(s.Appends) > 0)) || len(all) > p.MaxRuns || float64(s.Dead)/float64(total) > p.StaleRatio {
		t.Kind, t.Inputs = TaskMajor, all
		return t, true
	}
	if len(s.Appends) > 0 && float64(apps)/float64(total) > p.AppendRatio {
		t.Kind, t.Inputs = TaskFlush, s.Appends
		return t, true
	}
	if in := tierInputs(s.Runs, p.Fanout); len(in) > 0 {
		t.Kind, t.Inputs = TaskTier, in
		return t, true
	}
	return Task{}, false
}
