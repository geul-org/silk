//ff:func feature=index type=helper control=iteration dimension=1
//ff:what countOf — the records of one TAG kind in a document's counts
package index

// countOf finds tag in cs (0 when absent).
func countOf(cs []TagCount, tag uint16) uint32 {
	for _, c := range cs {
		if c.Tag == tag {
			return c.N
		}
	}
	return 0
}
