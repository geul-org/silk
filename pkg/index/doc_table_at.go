//ff:func feature=index type=decoder control=sequence
//ff:what DocTable.At — the entry of a document number
package index

import "sort"

func (t *DocTable) At(seq uint32) DocEntry {
	i := sort.Search(len(t.files), func(i int) bool { return t.files[i].first+t.files[i].count > seq })
	f := t.files[i]
	return parseDocEntry(f.raw[int(seq-f.first)*docSize:])
}
