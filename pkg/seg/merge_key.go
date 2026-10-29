//ff:func feature=seg type=builder control=sequence
//ff:what mergeKey — one partition from its groups: a streaming k-way merge when every input is sorted, else gathered and sorted in memory
package seg

// mergeKey writes groups of one key as one partition.
func mergeKey(gs []Group, w *Writer, live func(uint32) bool) (uint64, error) {
	if err := w.Begin(gs[0].Key); err != nil {
		return 0, err
	}
	var d uint64
	var err error
	if allSorted(gs) {
		d, err = kway(gs, w, live)
	} else {
		d, err = gather(gs, w, live)
	}
	w.End()
	return d, err
}
