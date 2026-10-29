//ff:func feature=recmap type=reader control=sequence
//ff:what File.Doc — the document holding a pack offset, by binary search over the document table
package recmap

import "sort"

// Doc returns the entry whose [Off, Off+Len) holds off.
func (f *File) Doc(off uint32) (DocEntry, bool) {
	i := sort.Search(len(f.Docs), func(i int) bool { return f.Docs[i].Off > off }) - 1
	if i < 0 || uint64(off) >= uint64(f.Docs[i].Off)+uint64(f.Docs[i].Len) {
		return DocEntry{}, false
	}
	return f.Docs[i], true
}
