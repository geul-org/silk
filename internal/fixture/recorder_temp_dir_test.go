//ff:func feature=fixture type=helper control=sequence
//ff:what recorder.TempDir — a directory of the real test
package fixture

func (r *recorder) TempDir() string { return r.TB.TempDir() }
