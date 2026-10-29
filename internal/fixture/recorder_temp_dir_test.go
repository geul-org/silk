//ff:func feature=fixture type=helper control=sequence
//ff:what recorder.TempDir — the preset directory when one is set (to make writes fail), else the real T's
package fixture

// TempDir lets a test hand Evolve an unusable directory.
func (r *recorder) TempDir() string {
	if r.tmp != "" {
		return r.tmp
	}
	return r.T.TempDir()
}
