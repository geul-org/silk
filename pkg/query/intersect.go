//ff:func feature=query type=builder control=iteration dimension=1
//ff:what intersect — documents in both the accumulated set and a condition's hits, Edge Ids appended; the first condition starts the set
package query

// intersect ANDs one condition's hits into acc.
func intersect(acc, hits map[docKey][]uint64, first bool) map[docKey][]uint64 {
	if first {
		return hits
	}
	out := map[docKey][]uint64{}
	for k, e := range acc {
		if h, ok := hits[k]; ok {
			out[k] = append(e, h...)
		}
	}
	return out
}
