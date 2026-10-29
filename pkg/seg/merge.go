//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what Merge — write sorted groups to a Writer key by key, dropping rows of dead documents; skip keys up to `from` (resume); call after() at each partition boundary
package seg

import "bytes"

// Merge returns the number of dead rows dropped. after may return an error
// to stop (for a checkpoint); the Writer is then between partitions.
func Merge(groups []Group, w *Writer, live func(uint32) bool, from []byte, after func(key []byte) error) (uint64, error) {
	dropped := uint64(0)
	for i := 0; i < len(groups); {
		j := keyEnd(groups, i)
		if from != nil && bytes.Compare(groups[i].Key, from) <= 0 {
			i = j
			continue
		}
		d, err := mergeKey(groups[i:j], w, live)
		dropped += d
		if err != nil {
			return dropped, err
		}
		if err := after(groups[i].Key); err != nil {
			return dropped, err
		}
		i = j
	}
	return dropped, nil
}
