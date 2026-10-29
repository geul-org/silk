//ff:func feature=index type=writer control=iteration dimension=1
//ff:what ingest.flushFull — write the buckets that reached the batch size
package index

// flushFull keeps memory bounded during a build.
func (g *ingest) flushFull() error {
	for i, b := range g.buckets {
		if b.Rows < g.batch {
			continue
		}
		if err := g.flushOne(i); err != nil {
			return err
		}
	}
	return nil
}
