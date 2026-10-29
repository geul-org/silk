//ff:func feature=seg type=helper control=iteration dimension=1
//ff:what keyEnd — the end of the run of groups sharing groups[i]'s key
package seg

import "bytes"

// keyEnd returns j > i with groups[i:j] of one key.
func keyEnd(groups []Group, i int) int {
	j := i + 1
	for j < len(groups) && bytes.Equal(groups[j].Key, groups[i].Key) {
		j++
	}
	return j
}
