//ff:func feature=seg type=helper control=iteration dimension=1
//ff:what minHead — index of the smallest non-nil row, -1 when every input is done
package seg

import "bytes"

// minHead picks the next row of a k-way merge (the first input wins ties).
func minHead(heads [][]byte) int {
	k := -1
	for i, h := range heads {
		if h != nil && (k < 0 || bytes.Compare(h, heads[k]) < 0) {
			k = i
		}
	}
	return k
}
