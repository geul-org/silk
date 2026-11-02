//ff:func feature=sig type=helper control=iteration dimension=1
//ff:what plan.product — the rows a plan spreads into
package sig

func (p plan) product() int {
	n := 1
	for _, c := range p.codes {
		n *= len(c)
	}
	return n
}
