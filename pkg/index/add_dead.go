//ff:func feature=index type=helper control=iteration dimension=1
//ff:what addDead — count a removed version's records as dead rows of every section that was not rebuilt
package index

// addDead adds counts' records of each section's TAG to its Dead.
func addDead(secs []Section, counts []TagCount, rebuild []bool) {
	for i := range secs {
		if !rebuild[i] {
			secs[i].Dead += uint64(countOf(counts, secs[i].Tag))
		}
	}
}
