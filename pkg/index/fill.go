//ff:func feature=index type=helper control=iteration dimension=1
//ff:what fill — a bool slice of n copies of v
package index

// fill makes a tag mask.
func fill(n int, v bool) []bool {
	out := make([]bool, n)
	for i := range out {
		out[i] = v
	}
	return out
}
