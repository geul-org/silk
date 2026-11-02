//ff:func feature=seg type=helper control=sequence
//ff:what Region.After — the first index at or after from whose leading bytes are above prefix
package seg

import (
	"bytes"
	"sort"
)

func (r Region) After(from int, prefix []byte) int {
	n := r.Len()
	return from + sort.Search(n-from, func(i int) bool { return bytes.Compare(r.Row(from + i)[:len(prefix)], prefix) > 0 })
}
