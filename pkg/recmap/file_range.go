//ff:func feature=recmap type=reader control=sequence
//ff:what File.Range — the key range of one Kind·Codec key by binary search
package recmap

import "sort"

// Range returns key's range; false when the map has no record of key.
func (f *File) Range(key uint16) (KeyRange, bool) {
	i := sort.Search(len(f.Keys), func(i int) bool { return f.Keys[i].Key >= key })
	if i == len(f.Keys) || f.Keys[i].Key != key {
		return KeyRange{}, false
	}
	return f.Keys[i], true
}
