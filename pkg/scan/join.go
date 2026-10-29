//ff:func feature=scan type=helper control=iteration dimension=1
//ff:what join — concatenate per-goroutine hits in chunk order; the first chunk error wins
package scan

// join merges the chunk results of Records.
func join(parts [][]Hit, errs []error) ([]Hit, error) {
	var out []Hit
	for i, p := range parts {
		if errs[i] != nil {
			return nil, errs[i]
		}
		out = append(out, p...)
	}
	return out, nil
}
