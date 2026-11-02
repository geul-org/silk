//ff:func feature=seg type=helper control=sequence
//ff:what Region.Seek — the first index at or after from whose row is not below key
package seg

import (
	"bytes"
	"sort"
)

func (r Region) Seek(from int, key []byte) int {
	n := r.Len()
	return from + sort.Search(n-from, func(i int) bool { return bytes.Compare(r.Row(from+i), key) >= 0 })
}
