//ff:func feature=cli type=helper control=sequence
//ff:what triples.Set — add one condition (flag.Value)
package cli

// Set appends v.
func (t *triples) Set(v string) error {
	*t = append(*t, v)
	return nil
}
