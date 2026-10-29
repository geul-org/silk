//ff:func feature=index type=helper control=iteration dimension=1
//ff:what addCount — count one record of a TAG kind in a document's (TAG, count) list
package index

// addCount increments tag's count, appending it when new (a document has
// few kinds, so a list beats a map).
func addCount(cs []TagCount, tag uint16) []TagCount {
	for i := range cs {
		if cs[i].Tag == tag {
			cs[i].N++
			return cs
		}
	}
	return append(cs, TagCount{Tag: tag, N: 1})
}
