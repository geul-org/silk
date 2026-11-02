//ff:func feature=index type=reader control=iteration dimension=1
//ff:what statSection — one section's counts
package index

import "slices"

func statSection(dir string, s Section, live Bitmap) (SectionStat, error) {
	d, err := readDecl(dir, s.Decl)
	if err != nil {
		return SectionStat{}, err
	}
	st := SectionStat{Name: s.Name, Kind: d.Kind, Bundle: s.Bundle, Pending: s.Pending, Runs: len(s.Runs), Appends: rowsOf(s.Appends)}
	st.Rows = rowsOf(s.Runs) + st.Appends
	for _, r := range append(slices.Clone(s.Runs), s.Appends...) {
		st.Bytes += r.Bytes
	}
	st.Stale, err = staleShare(dir, s, live)
	return st, err
}
