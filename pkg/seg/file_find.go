//ff:func feature=seg type=reader control=sequence
//ff:what File.Find — the partition table entry of a partition key slot, by binary search; false when the key is not its own partition
package seg

import (
	"bytes"
	"sort"
)

// Find returns the index in f.Parts.
func (f *File) Find(key []byte) (int, bool) {
	i := sort.Search(len(f.Parts), func(i int) bool { return bytes.Compare(f.Parts[i].Key, key) >= 0 })
	return i, i < len(f.Parts) && bytes.Equal(f.Parts[i].Key, key)
}
