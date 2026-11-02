//ff:func feature=index type=reader control=iteration dimension=1
//ff:what ReadStat — the current generation's counts
package index

func ReadStat(dir string) (Stat, error) {
	g, err := readCurrent(dir)
	if err != nil {
		return Stat{}, err
	}
	m, err := readManifest(dir, g)
	if err != nil {
		return Stat{}, err
	}
	live, err := loadLiveOf(dir, m)
	if err != nil {
		return Stat{}, err
	}
	st := Stat{Generation: g, Docs: m.DocCount(), Live: live.count(), MaxDoc: m.MaxDoc}
	for _, s := range m.Sections {
		ss, err := statSection(dir, s, live)
		if err != nil {
			return Stat{}, err
		}
		st.Sections = append(st.Sections, ss)
	}
	return st, nil
}
