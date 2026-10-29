//ff:func feature=index type=reader control=iteration dimension=1
//ff:what Index.Close — close every map and data pack file; the first error is returned
package index

// Close releases all parts.
func (ix *Index) Close() error {
	var first error
	for _, p := range ix.Parts {
		if err := p.Map.Close(); err != nil && first == nil {
			first = err
		}
		if err := p.Data.Close(); err != nil && first == nil {
			first = err
		}
	}
	return first
}
