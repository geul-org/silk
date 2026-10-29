//ff:func feature=index type=writer control=iteration dimension=1
//ff:what ingest.flush — write every non-empty bucket and close the data packs
package index

// flush finishes an ingest.
func (g *ingest) flush() error {
	for _, f := range g.files {
		f.Close()
	}
	g.files = nil
	for i := range g.buckets {
		if err := g.flushOne(i); err != nil {
			return err
		}
	}
	return nil
}
