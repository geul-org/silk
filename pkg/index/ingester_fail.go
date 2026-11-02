//ff:func feature=index type=writer control=sequence
//ff:what ingester.fail — remember the first error
package index

func (g *ingester) fail(err error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.err == nil {
		g.err = err
	}
	g.failed.Store(true)
}
