//ff:func feature=seg type=writer control=iteration dimension=1
//ff:what Bucket.Write — the gathered rows as a segment at path: partitions in key order, rows sorted (a run) or in arrival order (appended)
package seg

import (
	"slices"
)

// Write writes every partition; nothing is dropped.
func (b *Bucket) Write(path string, sorted bool) (Info, error) {
	info := b.Info
	info.Sorted = sorted
	w, err := NewWriter(path, info, b.Min)
	if err != nil {
		return Info{}, err
	}
	keys := make([]string, 0, len(b.parts))
	for k := range b.parts {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		if err := writePart(w, []byte(k), b.parts[k], sorted); err != nil {
			w.Abort()
			return Info{}, err
		}
	}
	return w.Finish()
}
