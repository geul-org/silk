//ff:func feature=index type=builder control=iteration dimension=1
//ff:what compactor.plan — the due merge of each section
package index

func (c *compactor) plan() ([]task, error) {
	var out []task
	secs := []*Section{&c.m.Paths}
	for i := range c.m.Sections {
		secs = append(secs, &c.m.Sections[i])
	}
	for _, s := range secs {
		t, ok, err := c.planSection(s)
		if err != nil {
			return nil, err
		}
		if ok {
			out = append(out, t)
		}
	}
	return out, nil
}
