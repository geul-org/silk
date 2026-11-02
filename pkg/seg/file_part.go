//ff:func feature=seg type=helper control=sequence
//ff:what File.Part — the rows of the partition with this key; false when there is none
package seg

import (
	"bytes"
	"sort"
)

func (f *File) Part(key []byte) (Region, bool) {
	i := uint64(sort.Search(int(f.Parts), func(i int) bool {
		k, _, _ := f.entry(uint64(i))
		return bytes.Compare(k, key) >= 0
	}))
	if i == f.Parts {
		return Region{W: f.RowWidth}, false
	}
	k, r := f.PartAt(i)
	if !bytes.Equal(k, key) {
		return Region{W: f.RowWidth}, false
	}
	return r, true
}
