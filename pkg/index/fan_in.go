//ff:func feature=index type=helper control=sequence
//ff:what fanIn — how many runs a merge may open under a memory limit
package index

func fanIn(limit int64) int {
	return int(min(max(limit/fanInBytes, minFanIn), maxFanIn))
}
