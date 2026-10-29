//ff:func feature=recmap type=builder control=iteration dimension=1
//ff:what sortByKey — counting sort of offsets by key (stable, so offsets stay ascending within a key) and the key range table
package recmap

// sortByKey sorts offs by keys in O(n). Offsets arrive ascending (documents
// in pack order, records in document order), so each key's range ascends.
func sortByKey(keys []uint16, offs []uint32) ([]KeyRange, []uint32) {
	var count [keySpace]uint64
	for _, k := range keys {
		count[k]++
	}
	var ranges []KeyRange
	var next [keySpace]uint64
	start := uint64(0)
	for k, c := range count {
		next[k] = start
		if c > 0 {
			ranges = append(ranges, KeyRange{Key: uint16(k), Start: start, Count: c})
		}
		start += c
	}
	sorted := make([]uint32, len(offs))
	for i, k := range keys {
		sorted[next[k]] = offs[i]
		next[k]++
	}
	return ranges, sorted
}
