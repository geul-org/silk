//ff:func feature=index type=builder control=iteration dimension=1
//ff:what updater.targets — group sections by declaration
package index

func (u *updater) targets() ([]*target, error) {
	g, err := loadGrammar(u.dir, u.m.Grammar)
	if err != nil {
		return nil, err
	}
	by := map[string]*target{}
	var out []*target
	for i := range u.m.Sections {
		s := &u.m.Sections[i]
		t, err := u.targetOf(by, s, g)
		if err != nil {
			return nil, err
		}
		t.sec = append(t.sec, s)
		if len(t.sec) == 1 {
			out = append(out, t)
		}
	}
	return out, nil
}
