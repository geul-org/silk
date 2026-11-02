//ff:func feature=index type=writer control=iteration dimension=1
//ff:what ingester.flushAll — sort every section's buffer into a run, a few at a time
package index

func (g *ingester) flushAll() error {
	var sps []*spill
	for _, t := range g.ts {
		sps = append(sps, t.sp...)
	}
	return parallel(len(sps), threadsOf(g.u.o), func(i int) error { return sps[i].flush() })
}
